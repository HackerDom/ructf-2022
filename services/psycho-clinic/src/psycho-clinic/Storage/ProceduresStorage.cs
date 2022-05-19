using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using psycho_clinic.Configuration;
using psycho_clinic.Extensions;
using psycho_clinic.Models;
using Vostok.Commons.Time;
using Vostok.Logging.Abstractions;

namespace psycho_clinic.Storage
{
    public class ProceduresStorage : IProceduresStorage
    {
        public ProceduresStorage(ISettingsProvider settingsProvider, ILog log)
        {
            this.settingsProvider = settingsProvider;
            dumpAction = new PeriodicalAction(() => Dump(), e => log.Error(e), () => 2.Seconds());

            dumpAction = new PeriodicalAction(
                () => Dump(),
                e => log.Error(e),
                () => settingsProvider.GetSettings().StorageDumpPeriod);
            dropAction = new PeriodicalAction(
                () => Drop(),
                e => log.Error(e),
                () => settingsProvider.GetSettings().StorageDropPeriod, true);
        }

        #region Service

        public void Start()
        {
            dumpAction.Start();
            dropAction.Start();
        }

        public void Stop()
        {
            dumpAction.Stop();
            dropAction.Stop();
        }

        public void Dump()
        {
            var dataPath = settingsProvider.GetSettings().ProceduresDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = proceduresByPatient
                .SelectMany(x => x.Value.Values)
                .ToArray();

            if (values.Length < 3)
                return;

            var tmpFileName = $"{dataPath}_tmp_{Guid.NewGuid()}";
            using (var tmpFile = new FileStream(tmpFileName, FileMode.Create))
            {
                tmpFile.Write(Encoding.UTF8.GetBytes(values.ToJson()));
            }

            File.Replace(tmpFileName, dataPath, null);
        }

        public void Drop()
        {
            proceduresByPatient.Clear();
            proceduresByPatient = new();

            File.Delete(settingsProvider.GetSettings().ProceduresDataPath);
        }

        public void Initialize(IEnumerable<TreatmentProcedure>? initialProcedures)
        {
            if (initialProcedures == null)
                return;

            foreach (var procedure in initialProcedures)
                AddProcedure(procedure.PatientId, procedure);
        }

        #endregion

        public List<TreatmentProcedure> GetPatientProcedures(PatientId patientId)
        {
            return proceduresByPatient.TryGetValue(patientId, out var procedures)
                ? procedures.Values.ToList()
                : new List<TreatmentProcedure>();
        }

        public bool GetPatientProcedure(
            PatientId patientId,
            TreatmentProcedureId procedureId,
            out TreatmentProcedure procedure
        )
        {
            procedure = null;

            return proceduresByPatient.TryGetValue(patientId, out var procedures) &&
                   procedures.TryGetValue(procedureId, out procedure);
        }

        public bool AddProcedure(PatientId patientId, TreatmentProcedure procedure)
        {
            var userProcedures = proceduresByPatient.GetOrAdd(patientId,
                _ => new ConcurrentDictionary<TreatmentProcedureId, TreatmentProcedure>());

            if (!userProcedures.TryAdd(procedure.Id, procedure))
                throw new Exception($"Procedure with id: {procedure.Id} already exists");

            return true;
        }

        private readonly PeriodicalAction dumpAction;
        private readonly PeriodicalAction dropAction;
        private readonly ISettingsProvider settingsProvider;

        private ConcurrentDictionary<PatientId, ConcurrentDictionary<TreatmentProcedureId, TreatmentProcedure>>
            proceduresByPatient = new();
    }
}
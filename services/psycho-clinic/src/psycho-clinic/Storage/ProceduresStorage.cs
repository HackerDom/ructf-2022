using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using psycho_clinic.Configuration;
using psycho_clinic.Extensions;
using psycho_clinic.Requests;
using Vostok.Commons.Time;
using Vostok.Logging.Abstractions;

namespace psycho_clinic.Storage
{
    public class ProceduresStorage : IProceduresStorage
    {
        public ProceduresStorage(ISettingsProvider settingsProvider, ILog log)
        {
            this.settingsProvider = settingsProvider;
            action = new PeriodicalAction(() => Dump(), e => log.Error(e), () => 2.Seconds());
        }

        #region Service

        public void Start()
        {
            action.Start();
        }

        public void Stop()
        {
            action.Stop();
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

            return userProcedures.TryAdd(procedure.Id, procedure);
        }

        private readonly PeriodicalAction action;
        private readonly ISettingsProvider settingsProvider;

        private readonly ConcurrentDictionary<PatientId, ConcurrentDictionary<TreatmentProcedureId, TreatmentProcedure>>
            proceduresByPatient = new();
    }
}
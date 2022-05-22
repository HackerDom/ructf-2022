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
            this.log = log.ForContext<ProceduresStorage>();

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

        public void Dump(bool allValues = false)
        {
            var dataPath = settingsProvider.GetSettings().ProceduresDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = proceduresByPatient
                .SelectMany(x => x.Value.Values)
                .ToArray();

            if (!allValues && values.Length < 3)
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
            if (!ClinicSettings.CleanerEnabled)
                return;

            log.Info("Starting to drop stale data");
            var expiredTime = DateTime.UtcNow - settingsProvider.GetSettings().StorageDataTTL;

            foreach (var (_, value) in proceduresByPatient)
            foreach (var (procedureId, timedValue) in value)
                if (timedValue.IsStale(expiredTime))
                {
                    value.Remove(procedureId, out _);
                    log.Info($"Removed {procedureId.Id}: {timedValue.TimeStamp}");
                }

            Dump(true);
        }

        public void Initialize(IEnumerable<TimedValue<TreatmentProcedure>>? initialProcedures)
        {
            if (initialProcedures == null)
                return;

            foreach (var timedValue in initialProcedures)
                AddProcedure(timedValue.Value.PatientId, timedValue.Value);
        }

        #endregion

        public List<TreatmentProcedure> GetPatientProcedures(PatientId patientId)
        {
            return proceduresByPatient.TryGetValue(patientId, out var procedures)
                ? procedures.Values.Select(x => x.Value).ToList()
                : new List<TreatmentProcedure>();
        }

        public bool GetPatientProcedure(
            PatientId patientId,
            TreatmentProcedureId procedureId,
            out TreatmentProcedure procedure
        )
        {
            procedure = null;

            if (!proceduresByPatient.TryGetValue(patientId, out var procedures))
                return false;

            if (!procedures.TryGetValue(procedureId, out var timedValue))
                return false;

            procedure = timedValue.Value;
            return true;
        }

        public bool AddProcedure(PatientId patientId, TreatmentProcedure procedure)
        {
            var userProcedures = proceduresByPatient.GetOrAdd(patientId,
                _ => new ConcurrentDictionary<TreatmentProcedureId, TimedValue<TreatmentProcedure>>());

            var procedureValue = new TimedValue<TreatmentProcedure>(procedure, DateTime.UtcNow);
            if (!userProcedures.TryAdd(procedure.Id, procedureValue))
                throw new Exception($"Procedure with id: {procedure.Id} already exists");

            return true;
        }

        public void Remove(PatientId patientId)
        {
            proceduresByPatient.Remove(patientId, out _);
        }

        private readonly PeriodicalAction dumpAction;
        private readonly PeriodicalAction dropAction;
        private readonly ISettingsProvider settingsProvider;
        private readonly ILog log;

        private readonly ConcurrentDictionary<PatientId,
                ConcurrentDictionary<TreatmentProcedureId, TimedValue<TreatmentProcedure>>>
            proceduresByPatient = new();
    }
}
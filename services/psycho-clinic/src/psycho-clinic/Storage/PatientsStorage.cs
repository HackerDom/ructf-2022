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
    public class PatientsStorage : IPatientsStorage
    {
        public PatientsStorage(
            ISettingsProvider settingsProvider,
            IContractsStorage contractsStorage,
            IProceduresStorage proceduresStorage,
            ILog log)
        {
            this.settingsProvider = settingsProvider;
            this.contractsStorage = contractsStorage;
            this.proceduresStorage = proceduresStorage;
            this.log = log.ForContext<PatientsStorage>();

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
            dumpAction.Stop();
        }

        public void Dump(bool allValues = false)
        {
            var dataPath = settingsProvider.GetSettings().PatientsDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = patients.Select(x => x.Value).ToArray();

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

            foreach (var (key, value) in patients)
                if (value.IsStale(expiredTime))
                {
                    patients.Remove(key, out _);
                    patientsByTokens.Remove(value.Value.Token, out _);
                    contractsStorage.Remove(key);
                    proceduresStorage.Remove(key);
                    log.Info($"Removed {key.Id}: {value.TimeStamp}");
                }

            Dump(true);
        }

        public void Initialize(IEnumerable<TimedValue<Patient>>? initialPatients)
        {
            if (initialPatients == null)
                return;

            foreach (var user in initialPatients)
                AddInternal(user);
        }

        #endregion

        public Patient Add(PatientId id, string name, DiagnosisType diagnosis)
        {
            patients.TryGetValue(id, out var patient);

            patient ??= new TimedValue<Patient>(new Patient(id, name, diagnosis), DateTime.Now);

            return AddInternal(patient);
        }

        public bool IsPatientExists(PatientToken patientToken, out Patient? patient)
        {
            return patientsByTokens.TryGetValue(patientToken, out patient);
        }

        private Patient AddInternal(TimedValue<Patient> patientValue)
        {
            var patient = patientValue.Value;
            try
            {
                if (!patients.TryAdd(patient.Id, patientValue))
                    throw new Exception($"Patient with id: {patient.Id} already exists.");
            }
            finally
            {
                patientsByTokens[patient.Token] = patient;
            }

            return patient;
        }

        private readonly PeriodicalAction dumpAction;
        private readonly PeriodicalAction dropAction;
        private readonly ISettingsProvider settingsProvider;
        private readonly IContractsStorage contractsStorage;
        private readonly IProceduresStorage proceduresStorage;
        private readonly ILog log;

        private readonly ConcurrentDictionary<PatientId, TimedValue<Patient>> patients = new();
        private readonly ConcurrentDictionary<PatientToken, Patient> patientsByTokens = new();
    }
}
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
        public PatientsStorage(ISettingsProvider settingsProvider, ILog log)
        {
            this.settingsProvider = settingsProvider;

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

        public void Dump()
        {
            var dataPath = settingsProvider.GetSettings().PatientsDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = patients.Select(x => x.Value).ToArray();

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
            patients.Clear();
            patientsByTokens.Clear();

            patients = new();
            patientsByTokens = new();

            File.Delete(settingsProvider.GetSettings().PatientsDataPath);
        }

        public void Initialize(IEnumerable<Patient>? initialPatients)
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

            patient ??= new Patient(id, name, diagnosis);

            return AddInternal(patient);
        }

        public bool IsPatientExists(PatientToken patientToken, out Patient? patient)
        {
            return patientsByTokens.TryGetValue(patientToken, out patient);
        }

        private Patient AddInternal(Patient patient)
        {
            try
            {
                if (!patients.TryAdd(patient.Id, patient))
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

        private ConcurrentDictionary<PatientId, Patient> patients = new();
        private ConcurrentDictionary<PatientToken, Patient> patientsByTokens = new();
    }
}
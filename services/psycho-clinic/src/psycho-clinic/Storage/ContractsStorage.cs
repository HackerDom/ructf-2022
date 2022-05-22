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
    public class ContractsStorage : IContractsStorage
    {
        public ContractsStorage(ISettingsProvider settingsProvider, ILog log)
        {
            this.settingsProvider = settingsProvider;
            this.log = log.ForContext<ContractsStorage>();

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
            var dataPath = settingsProvider.GetSettings().ContractsDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = contractsByPatient
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

            foreach (var (_, value) in contractsByPatient)
            foreach (var (contractId, timedValue) in value)
                if (timedValue.IsStale(expiredTime))
                {
                    value.Remove(contractId, out _);
                    log.Info($"Removed {contractId.Id}: {timedValue.TimeStamp}");
                }

            Dump(true);
        }

        public void Initialize(IEnumerable<TimedValue<Contract>>? initialContracts)
        {
            if (initialContracts == null)
                return;

            foreach (var timedValue in initialContracts)
                AddContract(timedValue.Value.Info.PatientId, timedValue.Value);
        }

        #endregion

        public List<Contract> GetPatientContracts(PatientId patientId)
        {
            return contractsByPatient.TryGetValue(patientId, out var contracts)
                ? contracts.Select(x => x.Value.Value).ToList()
                : new List<Contract>();
        }

        public bool AddContract(PatientId patientId, Contract contract)
        {
            var userContracts = contractsByPatient.GetOrAdd(patientId,
                _ => new ConcurrentDictionary<ContractId, TimedValue<Contract>>());

            var contractValue = new TimedValue<Contract>(contract, DateTime.UtcNow);
            if (!userContracts.TryAdd(contract.Id, contractValue))
                throw new Exception($"Contract with id: {contract.Id} already exists");

            return true;
        }

        public void Remove(PatientId patientId)
        {
            contractsByPatient.Remove(patientId, out _);
        }

        private readonly PeriodicalAction dumpAction;
        private readonly PeriodicalAction dropAction;
        private readonly ISettingsProvider settingsProvider;
        private readonly ILog log;

        private readonly ConcurrentDictionary<PatientId, ConcurrentDictionary<ContractId, TimedValue<Contract>>>
            contractsByPatient = new();
    }
}
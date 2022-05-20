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
            var dataPath = settingsProvider.GetSettings().ContractsDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = contractsByPatient
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
            contractsByPatient.Clear();
            contractsByPatient = new();

            File.Delete(settingsProvider.GetSettings().ContractsDataPath);
        }

        public void Initialize(IEnumerable<Contract>? initialContracts)
        {
            if (initialContracts == null)
                return;

            foreach (var contract in initialContracts)
                AddContract(contract.Info.PatientId, contract);
        }

        #endregion

        public List<Contract> GetPatientContracts(PatientId patientId)
        {
            return contractsByPatient.TryGetValue(patientId, out var contracts)
                ? contracts.Select(x => x.Value).ToList()
                : new List<Contract>();
        }

        public bool AddContract(PatientId patientId, Contract contract)
        {
            var userContracts = contractsByPatient.GetOrAdd(patientId,
                _ => new ConcurrentDictionary<ContractId, Contract>());

            if (!userContracts.TryAdd(contract.Id, contract))
                throw new Exception($"Contract with id: {contract.Id} already exists");

            return true;
        }

        private readonly PeriodicalAction dumpAction;
        private readonly PeriodicalAction dropAction;
        private readonly ISettingsProvider settingsProvider;

        private ConcurrentDictionary<PatientId, ConcurrentDictionary<ContractId, Contract>>
            contractsByPatient = new();
    }
}
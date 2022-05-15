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
                ? contracts.Values.ToList()
                : new List<Contract>();
        }

        /*public IEnumerable<Medicine> GetPossibleMedicines(int skip = 0, int take = 10)
        {
            if (take < 1)
                throw new ArgumentException("Parameter take should be greater then 0");
            if (skip < 0)
                throw new ArgumentException("Parameter skip should be greater or equal to 0");
            take = Math.Min(take, 100);

            //albumIds.Select(x => albums[x]).Select(x => new AlbumEntry(x.Id, x.Name, date)).ToArray();

            return medicines.Select(pair => pair.Value).ToArray();
        }*/

        public bool AddContract(PatientId patientId, Contract contract)
        {
            var userContracts = contractsByPatient.GetOrAdd(patientId,
                _ => new ConcurrentDictionary<ContractId, Contract>());

            return userContracts.TryAdd(contract.Id, contract);
        }

        private readonly PeriodicalAction action;
        private readonly ISettingsProvider settingsProvider;

        private readonly ConcurrentDictionary<PatientId, ConcurrentDictionary<ContractId, Contract>>
            contractsByPatient = new();
    }
}
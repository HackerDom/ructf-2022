using System.Collections.Generic;
using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IContractsStorage : IManagedStorage<Contract>
    {
        List<Contract> GetPatientContracts(PatientId patientId);
        bool AddContract(PatientId patientId, Contract contract);
    }
}
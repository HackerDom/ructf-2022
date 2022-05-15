using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IPatientsStorage : IManagedStorage<Patient>
    {
        Patient Add(PatientId id, string name, DiagnosisType diagnosis);

        bool IsPatientExists(PatientToken token, out Patient? patient);
    }
}
using psycho_clinic.Models;

namespace psycho_clinic.Requests
{
    public record RegisterPatientRequest(PatientId Id, string Name, DiagnosisType Diagnosis);
}
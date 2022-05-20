namespace psycho_clinic.Models
{
    public record TreatmentProcedureReport(
        TreatmentProcedureId ProcedureId,
        PatientId PatientId,
        DoctorId DoctorId,
        TreatmentProcedureResult ProcedureResult);
}
using psycho_clinic;

namespace psycho
{
    public record TreatmentProcedureReport(
        TreatmentProcedureId ProcedureId,
        PatientId PatientId,
        DoctorId DoctorId,
        TreatmentProcedureResult ProcedureResult);
}
using psycho_clinic.Models;

namespace psycho_clinic.Requests
{
    public record CreateReportRequest(TreatmentProcedureId ProcedureId, DoctorId DoctorId);
}
namespace psycho_clinic.Requests
{
    public record GetReportRequest(string DoctorName, string ProcedureId, int Skip);
}
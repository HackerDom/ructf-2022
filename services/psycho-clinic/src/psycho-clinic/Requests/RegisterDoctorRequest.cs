using psycho_clinic.Models;

namespace psycho_clinic.Requests
{
    public record RegisterDoctorRequest(DoctorId Id, string Name, string ProcedureDescription,
        EducationLevel EducationLevel);
}
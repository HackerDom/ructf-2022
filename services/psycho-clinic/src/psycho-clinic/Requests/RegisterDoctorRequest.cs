using psycho_clinic;
using psycho_clinic.Models;

namespace psycho
{
    public record RegisterDoctorRequest(DoctorId Id, string Name, string ProcedureDescription, EducationLevel EducationLevel);
}
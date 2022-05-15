using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IDoctorsStorage : IManagedStorage<Doctor>
    {
        bool Get(DoctorId doctorId, out Doctor doctor);
        Doctor Add(DoctorId doctorId, string name, string procedureDescription, EducationLevel educationLevel);
    }
}
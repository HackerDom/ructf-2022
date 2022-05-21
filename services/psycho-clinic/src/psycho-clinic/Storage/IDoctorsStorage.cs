using System.Collections.Generic;
using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IDoctorsStorage : IManagedStorage<Doctor>
    {
        bool TryGet(DoctorId doctorId, out Doctor doctor);
        Doctor Add(DoctorId doctorId, string name, string procedureDescription, EducationLevel educationLevel);
        IEnumerable<Doctor> GetDoctors();
    }
}
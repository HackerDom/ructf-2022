using System.Collections.Generic;

namespace psycho_clinic.Models
{
    public record GetDoctorsModel(int Count, IEnumerable<DoctorModel> Doctors);
}
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using psycho_clinic.Models;
using psycho_clinic.Requests;
using psycho_clinic.Storage;

namespace psycho_clinic.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class DoctorsController
    {
        public DoctorsController(IDoctorsStorage storage)
        {
            this.storage = storage;
        }

        [HttpPost("create/")]
        public Task<Doctor> Create(RegisterDoctorRequest request)
        {
            var (id, name, procedureDescription, educationLevel) = request;

            var doctor = storage.Add(id, name, procedureDescription, educationLevel);

            return Task.FromResult(doctor);
        }

        [HttpGet]
        public Task<IEnumerable<DoctorModel>> GetDoctors(int skip = 0)
        {
            var doctors = storage
                .GetDoctors()
                .Skip(skip)
                .Take(Take)
                .Select(d => new DoctorModel(d.Id, d.Name, d.EducationLevel));

            return Task.FromResult(doctors);
        }

        [HttpPost]
        public Task<DoctorModel> GetDoctor([FromBody] DoctorId doctorId)
        {
            if (!storage.TryGet(doctorId, out var doctor))
                throw new KeyNotFoundException($"Doctor with id {doctorId} was not found");

            return Task.FromResult(new DoctorModel(doctor.Id, doctor.Name, doctor.EducationLevel));
        }

        private const int Take = 10;
        private readonly IDoctorsStorage storage;
    }
}
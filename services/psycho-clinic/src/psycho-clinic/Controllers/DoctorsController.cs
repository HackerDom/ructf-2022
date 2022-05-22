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
        public Task<GetDoctorsModel> GetDoctors(EducationLevel educationLevel, int skip = 0, int take = 10)
        {
            var doctors = storage
                .GetDoctors()
                .Where(d => d.EducationLevel == educationLevel)
                .ToList();

            var count = doctors.Count;
            var result = doctors
                .Skip(skip)
                .Take(take)
                .Select(d => new DoctorModel(d.Id, d.Name, d.EducationLevel))
                .ToList();

            return Task.FromResult(new GetDoctorsModel(count, result));
        }

        [HttpPost]
        public Task<DoctorModel> GetDoctor([FromBody] DoctorId doctorId)
        {
            if (!storage.TryGet(doctorId, out var doctor))
                throw new KeyNotFoundException($"Doctor with id {doctorId} was not found");

            return Task.FromResult(new DoctorModel(doctor.Id, doctor.Name, doctor.EducationLevel));
        }

        private readonly IDoctorsStorage storage;
    }
}
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using psycho_clinic.Models;
using psycho_clinic.Requests;
using psycho_clinic.Storage;

namespace psycho_clinic.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class DoctorController
    {
        public DoctorController(IDoctorsStorage storage)
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

        [HttpPost]
        public Task<Doctor> GetDoctor([FromBody] DoctorId doctorId)
        {
            if (!storage.Get(doctorId, out var doctor))
                throw new KeyNotFoundException($"Doctor with id {doctorId} was not found");

            return Task.FromResult(doctor);
        }

        private readonly IDoctorsStorage storage;
    }
}
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using psycho_clinic.AppInfrastructure;
using psycho_clinic.Models;
using psycho_clinic.Requests;
using psycho_clinic.Storage;

namespace psycho_clinic.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class PatientController
    {
        public PatientController(IPatientsStorage storage)
        {
            this.storage = storage;
        }

        [HttpPost("create/")]
        public Task<PatientToken> Create(RegisterPatientRequest request)
        {
            var (id, name, diagnosisType) = request;

            var patient = storage.Add(id, name, diagnosisType);

            return Task.FromResult(patient.Token);
        }

        [HttpGet("card/")]
        public Task<Patient> GetPatientCard()
        {
            return Task.FromResult(Context.GetAuthenticatedPatient());
        }

        private readonly IPatientsStorage storage;
    }
}
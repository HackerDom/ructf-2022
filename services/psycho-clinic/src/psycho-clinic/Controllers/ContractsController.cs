using System;
using System.Collections.Generic;
using System.Linq;
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
    public class ContractsController
    {
        public ContractsController(IContractsStorage storage, IDoctorsStorage doctorsStorage)
        {
            this.storage = storage;
            this.doctorsStorage = doctorsStorage;
        }

        [HttpPost("create/")]
        public void CreateContract([FromBody] CreateContractRequest request)
        {
            var patient = Context.GetAuthenticatedPatient();

            var (contractId, doctorId, doctorSignature, dateTime) = request;

            CheckDoctor(doctorId, doctorSignature);

            var contract = new Contract(
                contractId,
                new ContractInfo(patient.Id, doctorId),
                dateTime
            );

            storage.AddContract(patient.Id, contract);
        }

        [HttpGet("all/")]
        public Task<IEnumerable<ContractModel>> GetPatientContracts()
        {
            var patient = Context.GetAuthenticatedPatient();

            var contracts = storage
                .GetPatientContracts(patient.Id)
                .Select(c => new ContractModel(c.Id, c.Info, c.Expired));

            return Task.FromResult(contracts);
        }

        private void CheckDoctor(DoctorId doctorId, DoctorSignature doctorSignature)
        {
            if (!doctorsStorage.TryGet(doctorId, out var doctor))
                throw new KeyNotFoundException($"Doctor with id: {doctorId} does not exist");

            if (!doctor.Signature.Equals(doctorSignature))
                throw new InvalidOperationException($"Doctor signature is invalid");
        }

        private readonly IContractsStorage storage;
        private readonly IDoctorsStorage doctorsStorage;
    }
}
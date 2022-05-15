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
        public ContractsController(IContractsStorage storage)
        {
            this.storage = storage;
        }

        [HttpPost("create/")]
        public void CreateContract([FromBody] CreateContractRequest request)
        {
            var patient = Context.GetAuthenticatedPatient();

            var (contractId, doctorId, doctorSignature, dateTime) = request; //TODO: check doctor signature
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

        private readonly IContractsStorage storage;
    }
}
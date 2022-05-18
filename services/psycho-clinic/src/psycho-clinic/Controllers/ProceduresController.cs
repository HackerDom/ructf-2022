using System;
using System.Collections.Generic;
using System.IO;
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
    public class ProceduresController
    {
        public ProceduresController(
            IProceduresStorage proceduresStorage,
            IContractsStorage contractsStorage,
            IDoctorsStorage doctorsStorage,
            IReportsStorage<TreatmentProcedureReport> reportsStorage)
        {
            this.proceduresStorage = proceduresStorage;
            this.contractsStorage = contractsStorage;
            this.doctorsStorage = doctorsStorage;
            this.reportsStorage = reportsStorage;
        }

        [HttpPost("prescribe/")]
        public void PrescribeProcedure([FromBody] PrescribeProcedureRequest request)
        {
            var patient = Context.GetAuthenticatedPatient();
            var (contractModel, procedureId, procedureType) = request;
            var contract = new Contract(contractModel.ContractId, contractModel.Info, contractModel.Expired);

            if (!HasContractWithDoctor(patient, contract))
                throw new InvalidOperationException("Unable to prescribe a procedure without active contract");

            var procedure = new TreatmentProcedure(procedureId, patient.Id, contract.Info.DoctorId, procedureType);

            proceduresStorage.AddProcedure(patient.Id, procedure);
        }

        [HttpGet("all/")]
        public Task<List<TreatmentProcedure>> GetPatientProcedures()
        {
            var patient = Context.GetAuthenticatedPatient();

            var contracts = proceduresStorage.GetPatientProcedures(patient.Id);

            return Task.FromResult(contracts);
        }

        [HttpPost("perform/")]
        public Task<TreatmentProcedureReport> Perform(TreatmentProcedureId procedureId)
        {
            var patient = Context.GetAuthenticatedPatient();

            var report = PerformProcedure(patient, procedureId);

            return Task.FromResult(report);
        }

        [HttpPost("report/create")]
        public void CreateReport(CreateReportRequest request)
        {
            var patient = Context.GetAuthenticatedPatient();

            var (procedureId, doctorId) = request;
            doctorsStorage.TryGet(doctorId, out var doctor);

            reportsStorage.Add(
                new TreatmentProcedureReport(
                    procedureId,
                    patient.Id,
                    doctorId,
                    new TreatmentProcedureResult(true, "asd")
                ),
                Path.Combine(doctor.Name, procedureId.Id.ToString())
            );
        }

        [HttpPost("report/")]
        public async Task<string> GetReport(GetReportRequest request)
        {
            var (doctorName, procedureId) = request;

            var report = await reportsStorage.Get(Path.Combine(doctorName, procedureId));

            var result = string.Join("", report);
            return await Task.FromResult(result);
        }

        private TreatmentProcedureReport PerformProcedure(Patient patient, TreatmentProcedureId procedureId)
        {
            if (!proceduresStorage.GetPatientProcedure(patient.Id, procedureId, out var procedure))
                throw new KeyNotFoundException(
                    $"Procedure with id: {procedureId.Id} for patient with id: {patient.Id} was not found");

            if (!doctorsStorage.TryGet(procedure.DoctorId, out var doctor))
                throw new KeyNotFoundException($"Doctor with id: {procedure.DoctorId} was not found");

            return new TreatmentProcedureReport(
                procedureId,
                patient.Id,
                doctor.Id,
                new TreatmentProcedureResult(true, doctor.ProcedureDescription));
        }

        private bool HasContractWithDoctor(Patient patient, Contract contract)
        {
            var now = DateTime.Now;

            var patientContracts = contractsStorage.GetPatientContracts(patient.Id);

            return now < contract.Expired && patientContracts.Contains(contract);
        }

        private readonly IProceduresStorage proceduresStorage;
        private readonly IContractsStorage contractsStorage;
        private readonly IDoctorsStorage doctorsStorage;
        private readonly IReportsStorage<TreatmentProcedureReport> reportsStorage;
    }
}
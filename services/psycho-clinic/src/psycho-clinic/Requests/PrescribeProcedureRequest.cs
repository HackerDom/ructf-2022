﻿namespace psycho_clinic.Requests
{
    public record PrescribeProcedureRequest(
        ContractModel ContractModel,
        TreatmentProcedureId ProcedureId,
        ProcedureType ProcedureType);
}
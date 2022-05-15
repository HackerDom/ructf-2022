using System;


namespace psycho_clinic.Requests
{
    public record CreateContractRequest(
        ContractId Id,
        DoctorId DoctorId,
        DoctorSignature DoctorSignature,
        DateTime Expired
    );
}
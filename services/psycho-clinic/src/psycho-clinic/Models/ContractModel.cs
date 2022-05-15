using System;

namespace psycho_clinic.Models
{
    public record ContractModel(ContractId ContractId, ContractInfo Info, DateTime Expired);
}
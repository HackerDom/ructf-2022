using System;

namespace psycho_clinic
{
    public record ContractModel(ContractId ContractId, ContractInfo Info, DateTime Expired);
}
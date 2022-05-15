using System;
using System.Text;
using psycho_clinic.Formatting;

namespace psycho_clinic.Models
{
    public record TreatmentProcedureReport(
        TreatmentProcedureId ProcedureId,
        PatientId PatientId,
        DoctorId DoctorId,
        TreatmentProcedureResult ProcedureResult) : ISerializable
    {
        public string Serialize()
        {
            return new StringBuilder()
                .Append("Procedure: ")
                .Append(ProcedureId.Id.ToString().AsSpan().Slice(0, 8))
                .Append(" ;")
                .Append("IsSuccessful: ")
                .Append(ProcedureResult.IsSuccessful)
                .Append(" ;")
                .AppendLine()
                .ToString();
        }
    }
}
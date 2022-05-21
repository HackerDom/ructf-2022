using System.Collections.Generic;
using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IProceduresStorage : IManagedStorage<TreatmentProcedure>
    {
        bool AddProcedure(PatientId patientId, TreatmentProcedure procedure);
        List<TreatmentProcedure> GetPatientProcedures(PatientId patientId);
        bool GetPatientProcedure(PatientId patientId, TreatmentProcedureId procedureId, out TreatmentProcedure procedure);
        void Remove(PatientId patientId);
    }
}
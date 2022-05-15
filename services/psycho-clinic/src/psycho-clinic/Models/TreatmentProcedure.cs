namespace psycho_clinic.Requests
{
    public class TreatmentProcedure
    {
        public TreatmentProcedureId Id { get; }
        public PatientId PatientId { get; }
        public DoctorId DoctorId { get; }
        public ProcedureType ProcedureType { get; }

        public TreatmentProcedure(
            TreatmentProcedureId id,
            PatientId patientId,
            DoctorId doctorId,
            ProcedureType procedureType)
        {
            Id = id;
            PatientId = patientId;
            DoctorId = doctorId;
            ProcedureType = procedureType;
        }

        #region EqualityMembers

        public bool Equals(TreatmentProcedure other)
        {
            return Id.Equals(other.Id) &&
                   PatientId.Equals(other.PatientId) &&
                   DoctorId.Equals(other.DoctorId) &&
                   ProcedureType.Equals(other.ProcedureType);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((TreatmentProcedure)obj);
        }

        public override int GetHashCode()
        {
            return Id.GetHashCode();
        }

        #endregion
    }
}
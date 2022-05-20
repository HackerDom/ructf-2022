using psycho_clinic.Formatting.Rendering;

namespace psycho_clinic.Models
{
    public class Doctor
    {
        public DoctorId Id { get; }
        public string Name { get; }
        public EducationLevel EducationLevel { get; }
        public string ProcedureDescription { get; }
        public DoctorSignature Signature { get; }

        public Doctor(DoctorId id, string name, string procedureDescription, EducationLevel educationLevel)
        {
            Id = id;
            Name = name;
            ProcedureDescription = procedureDescription;
            EducationLevel = educationLevel;
            Signature = new DoctorSignature(Renderer.Render($"{Id.Id}{Name}{EducationLevel}"));
        }

        #region EqualityMembers

        public bool Equals(Doctor other)
        {
            return Signature.Equals(other.Signature);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((Doctor)obj);
        }

        public override int GetHashCode()
        {
            return Signature.GetHashCode();
        }

        #endregion
    }
}
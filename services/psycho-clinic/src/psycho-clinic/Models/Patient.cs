using psycho_clinic.Formatting.Rendering;

namespace psycho_clinic.Models
{
    public class Patient
    {
        public PatientId Id { get; }
        public string Name { get; }
        public DiagnosisType Diagnosis { get; }
        public PatientToken Token { get; }

        public Patient(PatientId id, string name, DiagnosisType diagnosis)
        {
            Id = id;
            Name = name;
            Diagnosis = diagnosis;
            Token = new PatientToken(Renderer.Render($"{id.Id}:{Diagnosis}"));
        }

        #region EqualityMembers

        public bool Equals(Patient other)
        {
            return Id.Equals(other.Id) &&
                   Name.Equals(other.Name) &&
                   Token.Equals(other.Token);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((Patient)obj);
        }

        public override int GetHashCode()
        {
            return Token.GetHashCode();
        }

        #endregion
    }
}
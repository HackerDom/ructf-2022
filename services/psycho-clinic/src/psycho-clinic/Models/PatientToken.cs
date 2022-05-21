namespace psycho_clinic.Models
{
    public class PatientToken
    {
        public string Value { get; }

        public PatientToken(string value)
        {
            Value = value;
        }

        public bool Equals(PatientToken other)
        {
            return Value.Equals(other.Value);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((PatientToken)obj);
        }

        public override int GetHashCode()
        {
            return Value.GetHashCode();
        }
    }
}
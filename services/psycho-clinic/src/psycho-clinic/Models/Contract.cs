using System;
using psycho_clinic.Formatting.Rendering;

namespace psycho_clinic.Models
{
    public class Contract
    {
        public ContractId Id { get; }
        public ContractInfo Info { get; }
        private ContractIdentity Identity { get; }
        public DateTime Expired { get; }

        public Contract(ContractId id, ContractInfo info, DateTime expired)
        {
            Id = id;
            Info = info;
            Expired = expired;
            Identity = new ContractIdentity(Renderer.Render(
                $"{info.PatientId.Id.ToString()}{info.DoctorId.Id.ToString()}:{expired.Date.Day}{expired.Date.Month}"));
        }

        #region EqualityMembers

        public bool Equals(Contract other)
        {
            return Id.Equals(other.Id) && Identity.Equals(other.Identity);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((Contract)obj);
        }

        public override int GetHashCode()
        {
            return HashCode.Combine(Id, Identity);
        }

        #endregion
    }
}
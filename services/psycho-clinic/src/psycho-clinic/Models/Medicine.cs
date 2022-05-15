using System;

using psycho_clinic.Extensions;
using psycho_clinic.Formatting.Rendering;

namespace psycho_clinic.Models
{
    public class Medicine
    {
        public Guid Id { get; }
        public string Name { get; }
        public ContractInfo Info { get; }
        public ContractIdentity Identity { get; }
        public DateTime Expired { get; }

        public Medicine(Guid id, string name, ContractInfo info, DateTime expired)
        {
            Id = id;
            Name = name;
            Info = info;
            Expired = expired;
            Identity = new ContractIdentity(Renderer.Render($"{id.ToString()}:{expired.Date.Day}{expired.Date.Month}"));
        }

        #region EqualityMembers

        public bool Equals(Medicine other)
        {
            return Identity.Equals(other.Identity);
        }

        public override bool Equals(object? obj)
        {
            if (ReferenceEquals(null, obj)) return false;
            if (ReferenceEquals(this, obj)) return true;
            if (obj.GetType() != this.GetType()) return false;
            return Equals((Medicine)obj);
        }

        public override int GetHashCode()
        {
            return Identity.GetHashCode();
        }

        #endregion
    }
}
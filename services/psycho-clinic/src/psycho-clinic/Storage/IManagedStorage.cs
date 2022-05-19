using System.Collections.Generic;

namespace psycho_clinic.Storage
{
    public interface IManagedStorage<T> : IClearableStorage
    {
        void Initialize(IEnumerable<T> initialElements);
        void Dump();
    }
}
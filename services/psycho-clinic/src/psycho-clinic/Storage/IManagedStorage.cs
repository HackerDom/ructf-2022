using System.Collections.Generic;
using psycho_clinic.Models;

namespace psycho_clinic.Storage
{
    public interface IManagedStorage<T> : IClearableStorage
    {
        void Initialize(IEnumerable<TimedValue<T>> initialElements);
        void Dump(bool allValues = false);
    }
}
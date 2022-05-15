using System.Collections.Generic;

namespace psycho_clinic.Storage
{
    public interface IManagedStorage<T>
    {
        public void Start();
        public void Stop();
        public void Initialize(IEnumerable<T> initialElements);
        public void Dump();
    }
}
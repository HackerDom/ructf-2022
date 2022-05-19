using System.Threading.Tasks;

namespace psycho_clinic.Storage
{
    public interface IReportsStorage<T> : IClearableStorage
    {
        Task Add(T item, string filename);
        Task<string[]> Get(string filename);
    }
}
using System.Threading.Tasks;

namespace psycho_clinic.Storage
{
    public interface IReportsStorage<T>
    {
        Task Add(T item, string filename);
        Task<string[]> Get(string filename);
    }
}
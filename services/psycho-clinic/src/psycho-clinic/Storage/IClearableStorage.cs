namespace psycho_clinic.Storage
{
    public interface IClearableStorage
    {
        void Start();
        void Stop();
        void Drop();
    }
}
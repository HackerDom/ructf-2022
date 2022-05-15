namespace psycho_clinic.Configuration
{
    public interface ISettingsProvider
    {
        ClinicSettings GetSettings();
    }
}
namespace psycho_clinic.Configuration
{
    internal class SettingsProvider : ISettingsProvider
    {
        public SettingsProvider(Vostok.Configuration.Abstractions.IConfigurationProvider configProvider)
        {
            provider = configProvider;
        }

        public ClinicSettings GetSettings()
        {
            return provider.Get<ClinicSettings>();
        }

        private readonly Vostok.Configuration.Abstractions.IConfigurationProvider provider;
    }
}
namespace psycho_clinic.Configuration
{
    internal class SettingsProvider : ISettingsProvider
    {
        public SettingsProvider(Vostok.Configuration.Abstractions.IConfigurationProvider configProvider)
        {
            provider = configProvider;
        }

        public ServiceSettings GetSettings()
        {
            return provider.Get<ServiceSettings>();
        }

        private readonly Vostok.Configuration.Abstractions.IConfigurationProvider provider;
    }
}
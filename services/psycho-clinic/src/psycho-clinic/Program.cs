using System.IO;
using System.Threading.Tasks;
using psycho_clinic.Configuration;
using Vostok.Hosting;
using Vostok.Hosting.Setup;
using Vostok.Logging.File.Configuration;

namespace psycho_clinic
{
    public static class Program
    {
        public static async Task Main()
        {
            var appDataPath = ClinicSettings.AppPrefix;
            var storageDataPath = ClinicSettings.StorageDataPrefix;

            if (!Directory.Exists(appDataPath))
                Directory.CreateDirectory(appDataPath);
            if (!Directory.Exists(storageDataPath))
                Directory.CreateDirectory(storageDataPath);

            var host = new VostokHost(
                new VostokHostSettings(
                    new ClinicApplication(),
                    SetUpEnvironment)
            );

            await host.RunAsync();
        }

        private static void SetUpEnvironment(IVostokHostingEnvironmentBuilder builder)
        {
            builder
                .SetupApplicationIdentity(identityBuilder => identityBuilder
                    .SetEnvironment("environment")
                    .SetProject("project")
                    .SetApplication("psycho_clinic")
                    .SetInstance("0"))
                .SetupSystemMetrics(settings => settings.EnableProcessMetricsLogging = false)
                .SetupSystemMetrics(settings => settings.EnableHostMetricsLogging = false)
                .SetupSystemMetrics(settings => settings.EnableGcEventsLogging = false)
                .DisableHercules()
                .DisableServiceBeacon()
                .DisableClusterConfig()
                .SetupLog(logBuilder => logBuilder.SetupFileLog(fileLogBuilder =>
                    fileLogBuilder.CustomizeSettings(settings =>
                    {
                        settings.FilePath = Path.Combine(ClinicSettings.AppPrefix, "logs", "log");
                        settings.RollingStrategy.Type = RollingStrategyType.BySize;
                        settings.RollingStrategy.MaxSize = 100 * 1024;
                        settings.RollingStrategy.MaxFiles = 10;
                    })))
                .SetPort(18323);
        }
    }
}
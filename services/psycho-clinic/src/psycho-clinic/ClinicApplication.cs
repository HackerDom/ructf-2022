using System;
using System.IO;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using psycho_clinic.AppInfrastructure;
using psycho_clinic.Configuration;
using psycho_clinic.Extensions;
using psycho_clinic.Storage;
using Vostok.Applications.AspNetCore;
using Vostok.Applications.AspNetCore.Builders;
using Vostok.Hosting.Abstractions;
using Vostok.Hosting.Abstractions.Requirements;
using Vostok.Throttling.Config;

namespace psycho_clinic
{
    [RequiresConfiguration(typeof(ClinicSettings))]
    internal class ClinicApplication : VostokAspNetCoreApplication<Startup>
    {
        private IPatientsStorage patientsStorage;
        private IContractsStorage contractsStorage;
        private IProceduresStorage proceduresStorage;
        private IDoctorsStorage doctorsStorage;

        public override void Setup(IVostokAspNetCoreApplicationBuilder builder, IVostokHostingEnvironment environment)
        {
            base.Setup(builder, environment);

            builder.SetupThrottling(
                throttling =>
                {
                    throttling.UseEssentials(() => new ThrottlingEssentials());
                    throttling.DisableMetrics();
                });
        }

        public override Task WarmupAsync(IVostokHostingEnvironment environment, IServiceProvider serviceProvider)
        {
            var settings = serviceProvider.GetService<ISettingsProvider>()!.GetSettings();

            patientsStorage = serviceProvider.GetService<IPatientsStorage>()!;
            contractsStorage = serviceProvider.GetService<IContractsStorage>()!;
            proceduresStorage = serviceProvider.GetService<IProceduresStorage>()!;
            doctorsStorage = serviceProvider.GetService<IDoctorsStorage>()!;

            InitializeStorage(patientsStorage, settings.PatientsDataPath);
            InitializeStorage(contractsStorage, settings.ContractsDataPath);
            InitializeStorage(proceduresStorage, settings.ProceduresDataPath);
            InitializeStorage(doctorsStorage, settings.DoctorsDataPath);

            patientsStorage.Start();
            contractsStorage.Start();
            proceduresStorage.Start();
            doctorsStorage.Start();

            return base.WarmupAsync(environment, serviceProvider);
        }

        private static void InitializeStorage<T>(IManagedStorage<T> storage, string dataPath)
        {
            if (!File.Exists(dataPath))
                return;

            var bytes = File.ReadAllBytes(dataPath);
            if (bytes.Length == 0)
                return;

            var initialElements = Encoding.UTF8.GetString(bytes).FromJson<T[]>();
            storage.Initialize(initialElements);
        }

        public override void DoDispose()
        {
            patientsStorage.Stop();
            contractsStorage.Stop();
            proceduresStorage.Stop();
            doctorsStorage.Stop();

            base.DoDispose();
        }
    }
}
using Microsoft.Extensions.DependencyInjection;
using psycho_clinic.AppInfrastructure.Middleware;
using psycho_clinic.Configuration;
using psycho_clinic.Models;
using psycho_clinic.Storage;
using Vostok.Hosting.Abstractions;
using Vostok.Logging.Abstractions;

namespace psycho_clinic.AppInfrastructure
{
    public class DependencyResolution
    {
        public static void RegisterDependencies(
            IServiceCollection services,
            IVostokHostingEnvironment vostokHostingEnvironment)
        {
            services.AddSingleton<ILog>(_ => vostokHostingEnvironment.Log);
            services.AddSingleton<ISettingsProvider>(_ =>
                new SettingsProvider(vostokHostingEnvironment.ConfigurationProvider));

            services.AddSingleton<IAuthenticator, Authenticator>();
            services.AddSingleton<IPatientsStorage, PatientsStorage>();
            services.AddSingleton<IContractsStorage, ContractsStorage>();
            services.AddSingleton<IProceduresStorage, ProceduresStorage>();
            services.AddSingleton<IDoctorsStorage, DoctorsStorage>();
            services.AddSingleton<IReportsStorage<TreatmentProcedureReport>, ReportsStorage<TreatmentProcedureReport>>();
        }
    }
}
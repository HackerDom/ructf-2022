using System;
using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Newtonsoft.Json;
using psycho_clinic.AppInfrastructure.Middleware;
using psycho_clinic.Configuration;
using psycho_clinic.Extensions;
using psycho_clinic.Formatting.Rendering;
using psycho_clinic.Storage;
using Vostok.Context;
using Vostok.Hosting.Abstractions;
using Vostok.Logging.Abstractions;

namespace psycho_clinic.AppInfrastructure
{
    public class Startup
    {
        private readonly IVostokHostingEnvironment vostokHostingEnvironment;
        private readonly ILog log;

        public Startup(IConfiguration configuration)
        {
            vostokHostingEnvironment = FlowingContext.Globals.Get<IVostokHostingEnvironment>();
            if (vostokHostingEnvironment == null)
                throw new ArgumentNullException(nameof(vostokHostingEnvironment));

            log = vostokHostingEnvironment.Log;

            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services
                .AddControllers()
                .AddNewtonsoftJson(
                    options =>
                    {
                        foreach (var converter in JsonSerialization.Settings.Converters)
                            options.SerializerSettings.Converters.Add(converter);

                        options.SerializerSettings.NullValueHandling = NullValueHandling.Ignore;

                        options.SerializerSettings.Error += (_, args) =>
                        {
                            if (!args.ErrorContext.Handled)
                            {
                                var exception = args.ErrorContext.Error;
                                throw new ValidationException(
                                    $"Bad request: {exception.GetType().Name}: {exception.Message}", exception);
                            }
                        };
                    });

            services.AddSingleton<ILog>(_ => vostokHostingEnvironment.Log);
            services.AddSingleton<ISettingsProvider>(_ =>
                new SettingsProvider(vostokHostingEnvironment.ConfigurationProvider));
            
            services.AddSingleton<IAuthenticator, Authenticator>();
            services.AddSingleton<IPatientsStorage, PatientsStorage>();
            services.AddSingleton<IContractsStorage, ContractsStorage>();
            services.AddSingleton<IProceduresStorage, ProceduresStorage>();
            services.AddSingleton<IDoctorsStorage, DoctorsStorage>();

            //DependencyResolution.RegisterDependencies(services, vostokHostingEnvironment);
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app
                .UseHttpsRedirection()
                .UseRouting();

            app.UseMiddleware<ExceptionHandleMiddleware>();
            app.UseMiddleware<AuthenticationMiddleware>();

            app.UseEndpoints(endpoints => { endpoints.MapControllers(); });

            StartUpdating(app.ApplicationServices);
        }

        private static void StartUpdating(IServiceProvider serviceProvider)
        {
            //serviceProvider.GetService<ISpacePicturesUpdater>().Start();
        }
    }

    internal class Secret : Attribute { }
}
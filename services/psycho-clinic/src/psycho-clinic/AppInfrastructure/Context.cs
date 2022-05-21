using System;
using psycho_clinic.Models;
using Vostok.Context;

namespace psycho_clinic.AppInfrastructure
{
    public class Context
    {
        public static Patient? Patient
        {
            private get => FlowingContext.Globals.Get<Patient>();
            set => FlowingContext.Globals.Set(value);
        }

        public static Patient GetAuthenticatedPatient()
        {
            return Patient ?? throw new Exception("Patient is empty");
        }
    }
}
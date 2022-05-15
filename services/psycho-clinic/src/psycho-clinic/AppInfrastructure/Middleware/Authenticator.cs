using System;
using System.Linq;
using System.Security.Authentication;
using psycho_clinic.Models;
using psycho_clinic.Storage;

namespace psycho_clinic.AppInfrastructure.Middleware
{
    public class Authenticator : IAuthenticator
    {
        public Authenticator(IPatientsStorage patientsStorage)
        {
            this.patientsStorage = patientsStorage;
        }

        public Patient? Authenticate(string? patientTokenString)
        {
            if (patientTokenString == null)
                return null;

            var tokenParts = patientTokenString
                .Split(':')
                .Select(p => p.Replace("{", "").Replace("}", ""))
                .ToList();

            if (tokenParts.Count < 2 ||
                !Guid.TryParse(tokenParts[0], out _) ||
                !Enum.TryParse<DiagnosisType>(tokenParts[1], out _))
                throw new AuthenticationException("A valid API token was not specified with the request.");

            if (!patientsStorage.IsPatientExists(new (patientTokenString), out var patient))
                throw new AuthenticationException("Patient for provided Patient-token was not found.");

            return patient;
        }

        private readonly IPatientsStorage patientsStorage;
    }
}
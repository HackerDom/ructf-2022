using psycho_clinic.Models;

namespace psycho_clinic.AppInfrastructure.Middleware
{
    public interface IAuthenticator
    {
        Patient? Authenticate(string? apiTokenString);
    }
}
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;

namespace psycho_clinic.AppInfrastructure.Middleware
{
    public class AuthenticationMiddleware
    {
        public AuthenticationMiddleware(IAuthenticator authenticator, RequestDelegate next)
        {
            this.authenticator = authenticator;
            this.next = next;
        }

        public async Task Invoke(HttpContext context)
        {
            var apiTokenString = Extract(context.Request);

            Context.Patient = authenticator.Authenticate(apiTokenString);

            await next(context);
        }

        private static string? Extract(HttpRequest request)
        {
            var apiTokenString = FindInAuthorizationHeader(request);
            if (!string.IsNullOrEmpty(apiTokenString))
                return apiTokenString;

            return request.Cookies.TryGetValue(ApiTokenCookieName, out apiTokenString)
                ? apiTokenString
                : null;
        }

        private static string? FindInAuthorizationHeader(HttpRequest request)
        {
            if (!request.Headers.TryGetValue("Authorization", out var authorizationHeader))
                return null;

            var parts = authorizationHeader.ToString().Split(' ');

            return parts.Length <= 1 ? null : parts[1];
        }

        public const string ApiTokenCookieName = "Patient-Token";

        private readonly IAuthenticator authenticator;
        private readonly RequestDelegate next;
    }
}
using Microsoft.AspNetCore.Mvc;
using psycho_clinic.Configuration;

namespace psycho_clinic.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ServiceController
    {
        [HttpPost("cleaner-on/")]
        public void TurnCleanerOn()
        {
            ClinicSettings.CleanerEnabled = true;
        }

        [HttpPost("cleaner-off/")]
        public void TurnCleanerOff()
        {
            ClinicSettings.CleanerEnabled = false;
        }
    }
}
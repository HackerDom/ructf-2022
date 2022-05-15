using System;
using System.ComponentModel.DataAnnotations;
using System.IO;

namespace psycho_clinic.Configuration
{
    public class ServiceSettings
    {
        [Required] public string PathBase { get; }
        [Required] public Guid ServiceAdminApiKey { get; }

        public string PatientsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "patients_data");
        public string ProceduresDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "procedures_data");
        public string ContractsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "contracts_data");
        public string DoctorsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "doctors_data");

        public static readonly string AppPrefix = Path.Combine(Environment.CurrentDirectory, "data");
        public static readonly string StorageDataPrefix = Path.Combine(AppPrefix, "storage");
    }
}
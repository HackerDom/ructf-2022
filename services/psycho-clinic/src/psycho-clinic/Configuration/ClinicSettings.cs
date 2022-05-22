using System;
using System.IO;
using Vostok.Commons.Time;

namespace psycho_clinic.Configuration
{
    public class ClinicSettings
    {
        public TimeSpan StorageDumpPeriod = 2.Seconds();
        public TimeSpan StorageDropPeriod = 2.Hours();
        public TimeSpan StorageDataTTL = 2.Hours();

        public string PatientsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "patients_data");
        public string ProceduresDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "procedures_data");
        public string ContractsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "contracts_data");
        public string DoctorsDataPath = Path.Combine(AppPrefix, StorageDataPrefix, "doctors_data");

        public static readonly string AppPrefix = Path.Combine(Environment.CurrentDirectory, "data");
        public static readonly string StorageDataPrefix = Path.Combine(AppPrefix, "storage");

        public static bool CleanerEnabled = false;
    }
}
using System;
using System.IO;
using System.Threading.Tasks;
using psycho_clinic.Configuration;
using psycho_clinic.Formatting;

namespace psycho_clinic.Storage
{
    public class ReportsStorage<TValue> : IReportsStorage<TValue>
        where TValue : ISerializable
    {
        public ReportsStorage(ISettingsProvider settingsProvider)
        {
            folder = settingsProvider.GetSettings().ReportsDataPath;
            if (!Directory.Exists(folder))
                Directory.CreateDirectory(folder);
        }

        public async Task Add(TValue item, string filename)
        {
            var path = GetPath(filename);

            var directoryName = Path.GetDirectoryName(path);
            if (!Directory.Exists(directoryName))
                Directory.CreateDirectory(directoryName);

            await File.AppendAllTextAsync(path, item.Serialize());
        }

        public async Task<string[]> Get(string filename)
        {
            return await File.ReadAllLinesAsync(GetPath(filename));
        }

        private string GetPath(string name)
        {
            if (name.Contains(".."))
                throw new ArgumentException("File path should not contain \"..\"");

            var path = Path.Combine(folder, name);

            var storageDirectory = Path.GetFileName(ClinicSettings.StorageDataPrefix);
            if (path.Contains(storageDirectory, StringComparison.OrdinalIgnoreCase))
                throw new ArgumentException("File path should not contain path to storage");

            return path;
        }

        private readonly string folder;
    }
}
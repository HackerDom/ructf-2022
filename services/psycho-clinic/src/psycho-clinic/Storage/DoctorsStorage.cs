using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using psycho_clinic.Configuration;
using psycho_clinic.Extensions;
using psycho_clinic.Models;
using Vostok.Commons.Time;
using Vostok.Logging.Abstractions;

namespace psycho_clinic.Storage
{
    public class DoctorsStorage : IDoctorsStorage
    {
        public DoctorsStorage(ISettingsProvider settingsProvider, ILog log)
        {
            this.settingsProvider = settingsProvider;

            action = new PeriodicalAction(() => Dump(), e => log.Error(e), () => 2.Seconds());
        }

        #region Service

        public void Start()
        {
            action.Start();
        }

        public void Stop()
        {
            action.Stop();
        }

        public void Dump()
        {
            var dataPath = settingsProvider.GetSettings().DoctorsDataPath;

            if (!File.Exists(dataPath))
                File.Create(dataPath).Dispose();

            var values = doctors.Select(x => x.Value).ToArray();

            if (values.Length < 3)
                return;

            var tmpFileName = $"{dataPath}_tmp_{Guid.NewGuid()}";
            using (var tmpFile = new FileStream(tmpFileName, FileMode.Create))
            {
                tmpFile.Write(Encoding.UTF8.GetBytes(values.ToJson()));
            }

            File.Replace(tmpFileName, dataPath, null);
        }

        public void Initialize(IEnumerable<Doctor>? initialDoctors)
        {
            if (initialDoctors == null)
                return;

            foreach (var doctor in initialDoctors)
                AddInternal(doctor);
        }

        #endregion

        public Doctor Add(DoctorId doctorId, string name, string procedureDescription, EducationLevel educationLevel)
        {
            doctors.TryGetValue(doctorId, out var doctor);

            doctor ??= new Doctor(doctorId, name, procedureDescription, educationLevel);

            return AddInternal(doctor);
        }

        public IEnumerable<Doctor> GetDoctors()
        {
            return doctors.Select(x => x.Value);
        }

        public bool TryGet(DoctorId doctorId, out Doctor doctor)
        {
            return doctors.TryGetValue(doctorId, out doctor);
        }

        private Doctor AddInternal(Doctor doctor)
        {
            if (!doctors.TryAdd(doctor.Id, doctor))
                throw new Exception($"Doctor with id: {doctor.Id} already exists.");

            return doctor;
        }

        private readonly PeriodicalAction action;
        private readonly ISettingsProvider settingsProvider;

        private readonly ConcurrentDictionary<DoctorId, Doctor> doctors = new();
    }
}
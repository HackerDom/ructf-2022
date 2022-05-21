using System;

namespace psycho_clinic.Models
{
    public class TimedValue<T>
    {
        public T Value { get; }
        public DateTime TimeStamp { get; }

        public TimedValue(T value, DateTime timeStamp)
        {
            Value = value;
            TimeStamp = timeStamp;
        }

        public bool IsStale(DateTime expiration)
        {
            return TimeStamp <= expiration;
        }
    }
}
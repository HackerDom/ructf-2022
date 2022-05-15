using System.Collections.Generic;

namespace psycho_clinic.Formatting
{
    public class PrintingContext
    {
        public PrintingContext(string messageTemplate)
            : this(messageTemplate, null)
        {
        }

        public PrintingContext(string messageTemplate, IReadOnlyDictionary<string, object> properties)
        {
            MessageTemplate = messageTemplate;
            Properties = properties;
        }

        public string MessageTemplate { get; }

        public IReadOnlyDictionary<string, object> Properties { get; }
    }
}
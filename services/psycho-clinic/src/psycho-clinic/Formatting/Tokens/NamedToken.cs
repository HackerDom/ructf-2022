using System;
using System.Text;

namespace psycho_clinic.Formatting.Tokens
{
    internal abstract class NamedToken : ITemplateToken
    {
        protected NamedToken(string name, string format)
        {
            Name = name ?? throw new ArgumentNullException(nameof(name));
            Format = format;
        }

        public string Name { get; }
        public string Format { get; }

        public abstract void Render(PrintingContext context, StringBuilder output);

        public override string ToString() => Format == null
            ? $"{{{Name}}}"
            : $"{{{Name}:{Format}}}";
    }
}
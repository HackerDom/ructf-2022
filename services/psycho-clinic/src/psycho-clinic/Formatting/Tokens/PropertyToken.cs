using System.Diagnostics.CodeAnalysis;
using System.Text;

namespace psycho_clinic.Formatting.Tokens
{
    internal class PropertyToken : NamedToken
    {
        public PropertyToken([NotNull] string name, string format = null)
            : base(name, format)
        {
        }

        public override void Render(PrintingContext context, StringBuilder output)
        {
            if (context.Properties == null)
                return;

            if (context.Properties.TryGetValue(Name, out var value))
            {
                output.Append($"{{{value}}}");
            }
        }
    }
}
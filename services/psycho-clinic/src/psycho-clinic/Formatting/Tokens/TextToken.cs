using System;
using System.Diagnostics.CodeAnalysis;
using System.Text;

namespace psycho_clinic.Formatting.Tokens
{
    internal class TextToken : ITemplateToken
    {
        public TextToken([NotNull] string text, int offset, int length)
        {
            Text = (text ?? throw new ArgumentNullException(nameof(text))).Substring(offset, length);
        }

        public TextToken([NotNull] string text)
        {
            Text = text ?? throw new ArgumentNullException(nameof(text));
        }

        public string Text { get; }

        public override string ToString() => Text;

        public void Render(PrintingContext context, StringBuilder output)
        {
            output.Append(Text);
        }
    }
}
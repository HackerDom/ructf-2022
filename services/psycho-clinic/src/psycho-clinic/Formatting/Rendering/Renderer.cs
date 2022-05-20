using System;
using System.Collections.Generic;
using System.Globalization;
using System.Runtime.CompilerServices;
using System.Text;
using psycho_clinic.Formatting.TokensFactory;

namespace psycho_clinic.Formatting.Rendering
{
    public static class Renderer
    {
        public static string Render(PrintingContext context)
        {
            var tokens = Tokenizer.Tokenize(context.MessageTemplate, new PropertyTokensFactory());
            var builder = new StringBuilder();

            foreach (var token in tokens)
                token.Render(context, builder);

            return builder.ToString();
        }

        public static string Render(ref InfoStringHandler message)
        {
            return Render(new PrintingContext(message.MessageTemplate.ToString(), message.Properties));
        }

        public static string Render(string template)
        {
            return Render(new PrintingContext(template));
        }

        [InterpolatedStringHandler]
        public ref struct InfoStringHandler
        {
            public InfoStringHandler(int literalLength, int formattedCount)
            {
                MessageTemplate = null;
                Properties = null;

                MessageTemplate = new StringBuilder(literalLength);

                if (formattedCount > 0)
                    Properties = new Dictionary<string, object>(StringComparer.Ordinal);
            }

            public void AppendLiteral(string value)
            {
                MessageTemplate.Append(value);
            }

            public void AppendFormatted<T>(T value, int alignment, [CallerArgumentExpression("value")] string name = "")
            {
                var defaultHandler = CreateDefaultHandler();
                defaultHandler.AppendFormatted(value, alignment);
                AppendFormatted((object)defaultHandler.ToStringAndClear(), name);
            }

            public void AppendFormatted<T>(T value, string format, [CallerArgumentExpression("value")] string name = "")
            {
                var defaultHandler = CreateDefaultHandler();
                defaultHandler.AppendFormatted(value, format);
                AppendFormatted((object)defaultHandler.ToStringAndClear(), name);
            }

            public void AppendFormatted<T>(T value, int alignment, string format, [CallerArgumentExpression("value")] string name = "")
            {
                var defaultHandler = CreateDefaultHandler();
                defaultHandler.AppendFormatted(value, alignment, format);
                AppendFormatted((object)defaultHandler.ToStringAndClear(), name);
            }

            public void AppendFormatted(object value, [CallerArgumentExpression("value")] string name = "")
            {
                MessageTemplate.Append('{');
                MessageTemplate.Append(name);
                MessageTemplate.Append('}');

                Properties.Add(name, value);
            }

            internal StringBuilder MessageTemplate { get; }
            internal Dictionary<string, object> Properties { get; private set; }

            private static DefaultInterpolatedStringHandler CreateDefaultHandler() =>
                new(0, 1, CultureInfo.InvariantCulture);
        }
    }
}
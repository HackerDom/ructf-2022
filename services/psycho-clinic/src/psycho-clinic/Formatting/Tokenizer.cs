using System.Collections.Generic;
using System.Runtime.CompilerServices;
using System.Text;
using psycho_clinic.Formatting.Tokens;
using psycho_clinic.Formatting.TokensFactory;

namespace psycho_clinic.Formatting
{
    internal static class Tokenizer
    {
        private static readonly char[] AllBraces = { OpeningBrace, ClosingBrace };

        private const char OpeningBrace = '{';
        private const char ClosingBrace = '}';
        private const char Underscore = '_';
        private const char Whitespace = ' ';
        private const char Colon = ':';
        private const char At = '@';
        private const char Dot = '.';

        public static IEnumerable<ITemplateToken> Tokenize(string template, INamedTokenFactory namedTokenFactory)
        {
            if (string.IsNullOrEmpty(template))
                yield break;

            var nextIndex = 0;

            while (true)
            {
                var text = ConsumeText(template, nextIndex, out nextIndex);
                if (text != null)
                    yield return text;

                if (nextIndex == template.Length)
                    yield break;

                var namedToken = ParseNamedToken(template, nextIndex, out nextIndex, namedTokenFactory);
                if (namedToken != null)
                    yield return namedToken;

                if (nextIndex == template.Length)
                    yield break;
            }
        }

        private static ITemplateToken ConsumeText(string template, int offset, out int next)
        {
            var beginning = offset;
            var builder = new StringBuilder();

            do
            {
                var currentCharacter = template[offset];
                if (currentCharacter == OpeningBrace)
                {
                    if (offset + 1 < template.Length && template[offset + 1] == OpeningBrace)
                    {
                        builder.Append(currentCharacter);
                        offset++;
                    }
                    else break;
                }
                else
                {
                    builder.Append(currentCharacter);

                    if (currentCharacter == ClosingBrace)
                    {
                        if (offset + 1 < template.Length && template[offset + 1] == ClosingBrace)
                        {
                            offset++;
                        }
                    }
                }

                offset++;
            } while (offset < template.Length);

            next = offset;

            return CreateTextToken(template, beginning, offset - beginning, builder);
        }

        private static ITemplateToken ParseNamedToken(
            string template,
            int offset,
            out int next,
            INamedTokenFactory factory)
        {
            var beginning = offset++;

            while (offset < template.Length && IsValidInNamedToken(template[offset]))
                offset++;

            if (offset == template.Length || template[offset] != ClosingBrace)
            {
                next = offset;

                return CreateTextToken(template, beginning, offset - beginning);
            }

            next = offset + 1;

            var rawOffset = beginning;
            var rawLength = next - rawOffset;
            var tokenOffset = rawOffset + 1;
            var tokenLength = rawLength - 2;

            if (TryParseNamedToken(template, tokenOffset, tokenLength, out var name, out var format))
                return factory.Create(name, format);

            return CreateTextToken(template, rawOffset, rawLength);
        }

        private static TextToken CreateTextToken(string template, int offset, int length)
        {
            return length == 0 ? null : new TextToken(template, offset, length);
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        private static TextToken CreateTextToken(string template, int offset, int length, StringBuilder builder)
        {
            if (length == 0)
                return null;

            if (length == builder.Length)
                return new TextToken(template, offset, length);

            return new TextToken(builder.ToString());
        }

        private static bool TryParseNamedToken(
            string template,
            int offset,
            int length,
            out string name,
            out string format)
        {
            if (length == 0)
            {
                name = format = null;
                return false;
            }

            var formatDelimiter = template.IndexOf(Colon, offset, length);
            if (formatDelimiter < offset)
            {
                name = template.Substring(offset, length);
                format = null;
            }
            else
            {
                name = template.Substring(offset, formatDelimiter - offset);
                format = template.Substring(formatDelimiter + 1, offset + length - formatDelimiter - 1);
            }

            if (string.IsNullOrEmpty(format))
                format = null;

            return IsValidName(name) && IsValidFormat(format);
        }

        private static bool IsValidName(string name)
        {
            if (string.IsNullOrEmpty(name))
                return false;

            foreach (var symbol in name)
                if (!IsValidInName(symbol))
                    return false;

            return true;
        }

        private static bool IsValidFormat(string format)
        {
            if (format == null)
                return true;

            foreach (var symbol in format)
                if (!IsValidInFormat(symbol))
                    return false;

            return true;
        }

        private static bool IsValidInNamedToken(char c) =>
            IsValidInName(c) || IsValidInFormat(c) || c == Colon;

        private static bool IsValidInName(char c) =>
            char.IsLetterOrDigit(c) || c == Underscore || c == At || c == Dot;

        private static bool IsValidInFormat(char c) =>
            c != ClosingBrace && (char.IsLetterOrDigit(c) || char.IsPunctuation(c) || c == Whitespace);
    }
}
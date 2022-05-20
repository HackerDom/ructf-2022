using psycho_clinic.Formatting.Tokens;

namespace psycho_clinic.Formatting.TokensFactory
{
    internal class PropertyTokensFactory : INamedTokenFactory
    {
        public ITemplateToken Create(string name, string format) => new PropertyToken(name, format);
    }
}
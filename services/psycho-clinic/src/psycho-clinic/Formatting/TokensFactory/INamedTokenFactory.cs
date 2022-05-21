using psycho_clinic.Formatting.Tokens;

namespace psycho_clinic.Formatting.TokensFactory
{
    internal interface INamedTokenFactory
    {
        ITemplateToken Create(string name, string format);
    }
}
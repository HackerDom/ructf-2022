using System.Text;

namespace psycho_clinic.Formatting.Tokens
{
    public interface ITemplateToken
    {
        void Render(PrintingContext context, StringBuilder output);
    }
}
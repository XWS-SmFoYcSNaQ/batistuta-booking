using FluentValidation;
using user_service.domain.Enums;

namespace user_service.Validators
{
    public class RegisterUserRequestValidator : AbstractValidator<RegisterUser_Request>
    {
        private const string _FirstCaseCapitalRegex = "^[A-Z]{1}[a-zA-Z]*";

        public RegisterUserRequestValidator()
        {
            RuleFor(x => x.Email)
                .EmailAddress();
            RuleFor(x => x.Password)
                .NotEmpty()
                .MinimumLength(8);
            RuleFor(x => x.Username)
                .NotEmpty();
            RuleFor(x => x.FirstName)
                .Matches(_FirstCaseCapitalRegex);
            RuleFor(x => x.LastName)
                .Matches(_FirstCaseCapitalRegex);
        }
    }
}

using FluentValidation;

namespace user_service.Validators
{
    public class RegisterUserRequestValidator : AbstractValidator<RegisterUser_Request>
    {
        private const string _firstCaseCapitalPattern = "^[A-Z]{1}[a-zA-Z]*";
        private const string _passwordPattern = "[0-9a-zA-Z]+";

        public RegisterUserRequestValidator()
        {
            RuleFor(x => x.Email)
                .EmailAddress();
            RuleFor(x => x.Password)
                .MinimumLength(8)
                .Matches(_passwordPattern);
            RuleFor(x => x.Username)
                .NotEmpty();
            RuleFor(x => x.FirstName)
                .Matches(_firstCaseCapitalPattern);
            RuleFor(x => x.LastName)
                .Matches(_firstCaseCapitalPattern);
        }
    }
}

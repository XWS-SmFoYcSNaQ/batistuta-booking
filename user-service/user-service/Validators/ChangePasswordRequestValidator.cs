using FluentValidation;

namespace user_service.Validators
{
    public class ChangePasswordRequestValidator : AbstractValidator<ChangePassword_Request>
    {
        private const string _passwordPattern = "[0-9a-zA-Z]+";

        public ChangePasswordRequestValidator()
        {
            RuleFor(x => x.NewPassword)
                .MaximumLength(8)
                .Matches(_passwordPattern);
        }
    }
}

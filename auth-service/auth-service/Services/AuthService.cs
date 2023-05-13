using Grpc.Core;

namespace auth_service.Services
{
    public class AuthService : auth_service.AuthService.AuthServiceBase
    {
        private readonly ILogger<AuthService> _logger;

        public AuthService(ILogger<AuthService> logger)
        {
            _logger = logger;
        }

        public override Task<VerifyResponse> Verify(VerifyRequest request, ServerCallContext context)
        {
            _logger.LogInformation("Verify triggered.");
            return Task.FromResult(new VerifyResponse
            {
                Message = request.Name + " verified"
            });
        }
    }
}

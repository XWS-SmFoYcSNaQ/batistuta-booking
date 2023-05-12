using Grpc.Core;

namespace user_service.Services
{
    public class UserService : user_service.UserService.UserServiceBase
    {
        private readonly ILogger<UserService> _logger;

        public UserService(ILogger<UserService> logger)
        {
            _logger = logger;
        }

        public override Task<US_Response> GetUser(US_Request request, ServerCallContext context)
        {
            _logger.LogInformation("GetUser triggered.");
            return Task.FromResult(new US_Response
            {
                Message = $"Hello {request.Name}"
            });
        }
    }
}

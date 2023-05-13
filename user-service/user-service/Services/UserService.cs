using Grpc.Core;
using user_service.data.Db;

namespace user_service.Services
{
    public class UserService : user_service.UserService.UserServiceBase
    {
        private readonly ILogger<UserService> _logger;
        private readonly UserServiceDbContext _dbContext;

        public UserService(ILogger<UserService> logger, UserServiceDbContext dbContext)
        {
            _logger = logger;
            _dbContext = dbContext;
        }

        public override Task<US_Response> GetUser(US_Request request, ServerCallContext context)
        {
            _logger.LogInformation("GetUser triggered.");
            Console.WriteLine(_dbContext.Users.Count());
            return Task.FromResult(new US_Response
            {
                Message = $"Hello {request.Name}"
            });
        }
    }
}

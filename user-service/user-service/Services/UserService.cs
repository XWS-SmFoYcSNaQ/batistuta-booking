using AutoMapper;
using FluentValidation;
using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using user_service.data.Db;
using user_service.Interfaces;

namespace user_service.Services
{
    public class UserService : user_service.UserService.UserServiceBase
    {
        private readonly ILogger<UserService> _logger;
        private readonly UserServiceDbContext _dbContext;
        private readonly IMapper _mapper;
        private readonly IValidator<RegisterUser_Request> _validator;
        private readonly IPasswordHasher _passwordHasher;

        public UserService(ILogger<UserService> logger,
            UserServiceDbContext dbContext,
            IMapper mapper,
            IValidator<RegisterUser_Request> validator,
            IPasswordHasher passwordHasher)
        {
            _logger = logger;
            _dbContext = dbContext;
            _mapper = mapper;
            _validator = validator;
            _passwordHasher = passwordHasher;
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

        public override async Task<RegisterUser_Response> RegisterUser(RegisterUser_Request request, ServerCallContext context)
        {
            var validationResult = await _validator.ValidateAsync(request);

            if (!validationResult.IsValid)
            {
                var errorResponse = new RegisterUser_Response
                {
                    Message = "Error not all fields are valid."
                };
                errorResponse.Errors.AddRange(validationResult.Errors.Select(x => new RegisterUser_Response.Types.Error
                {
                    PropertyName = x.PropertyName,
                    ErrorMessage = x.ErrorMessage
                }));

                return errorResponse;
            }

            var user = _mapper.Map<user_service.Models.User>(request);
            var userEntity = _mapper.Map<user_service.domain.Entities.User>(user);
            userEntity.Password = _passwordHasher.Hash(userEntity.Password);
            _dbContext.Users.Add(userEntity);
            try
            {
                int rowsAffected = await _dbContext.SaveChangesAsync();
                if (rowsAffected > 0)
                {
                    return new RegisterUser_Response
                    {
                        Success = true,
                        Message = "User successfully registred"
                    };
                }
                return new RegisterUser_Response
                {
                    Message = "An unknown error occured while trying to register user."
                };
            }
            catch (Exception ex)
            {
                _logger.LogError(ex.ToString());
                return new RegisterUser_Response
                {
                    Message = ex.InnerException?.Message ?? ex.Message
                };
            }
        }

        public override async Task<GetAllUsers_Response> GetAllUsers(Empty_Request request, ServerCallContext context)
        {
            var users = await _dbContext.Users.ToListAsync();
            var response = new GetAllUsers_Response();

            response.Users.AddRange(users.Select(x => _mapper.Map<user_service.GetAllUsers_Response.Types.User>(x)));

            return response;
        }
    }
}

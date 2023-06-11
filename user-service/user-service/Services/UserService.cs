using AutoMapper;
using FluentValidation;
using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using user_service.Configuration;
using user_service.data.Db;
using user_service.Helpers;
using user_service.Interfaces;
using AuthServiceClient;
using Microsoft.AspNetCore.HttpOverrides;

namespace user_service.Services
{
    public class UserService : user_service.UserService.UserServiceBase
    {
        private readonly UserServiceDbContext _dbContext;
        private readonly IMapper _mapper;
        private readonly ILogger<UserService> _logger;
        private readonly IValidator<RegisterUser_Request> _validator;
        private readonly IPasswordHasher _passwordHasher;
        private readonly GrpcChannelBuilder _grpcChannelBuilder;
        private readonly ServicesConfig _servicesConfig;


        public UserService(UserServiceDbContext dbContext,
            IMapper mapper,
            ILogger<UserService> logger,
            IValidator<RegisterUser_Request> validator,
            IPasswordHasher passwordHasher,
            GrpcChannelBuilder grpcChannelBuilder,
            ServicesConfig servicesConfig)
        {
            _dbContext = dbContext;
            _mapper = mapper;
            _logger = logger;
            _validator = validator;
            _passwordHasher = passwordHasher;
            _grpcChannelBuilder = grpcChannelBuilder;
            _servicesConfig = servicesConfig;
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
                errorResponse.Errors.AddRange(validationResult.Errors.Select(x => new Error
                {
                    PropertyName = x.PropertyName,
                    ErrorMessage = x.ErrorMessage
                }));

                context.Status = new Status(StatusCode.InvalidArgument, "Not all fields are valid.");
                return errorResponse;
            }

            var user = _mapper.Map<Models.User>(request);
            var userEntity = _mapper.Map<domain.Entities.User>(user);
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
                        Message = "User successfully registred",
                        User = _mapper.Map<User>(userEntity)
                    };
                }

                context.Status = new Status(StatusCode.Internal, "An unknown error occured while trying to register user.");
                return new RegisterUser_Response
                {
                    Message = "An unknown error occured while trying to register user."
                };
            }
            catch (Exception ex)
            {
                if (ex is RpcException rpcException)
                {
                    throw new RpcException(rpcException.Status, rpcException.Trailers);
                }
                throw new RpcException(new Status(StatusCode.Internal, ex.InnerException?.Message ?? ex.Message));
            }
        }

        public override async Task<GetAllUsers_Response> GetAllUsers(Empty_Request request, ServerCallContext context)
        {
            try
            {
                var users = await _dbContext.Users.ToListAsync();
                var response = new GetAllUsers_Response();

                response.Users.AddRange(users.Select(x => _mapper.Map<UserLessInfo>(x)));

                return response;
            }
            catch (Exception ex)
            {
                if (ex is RpcException rpcException)
                {
                    throw new RpcException(rpcException.Status, rpcException.Trailers);
                }
                throw new RpcException(new Status(StatusCode.Internal, ex.InnerException?.Message ?? ex.Message), Metadata.Empty);
            }

        }

        public override async Task<VerifyUser_Response> VerifyUserPassword(VerifyUser_Request request, ServerCallContext context)
        {
            var username = request.Username;
            var password = request.Password;
            var user = await _dbContext.Users
                .Where(x => x.Username == username)
                .FirstOrDefaultAsync();

            if (user == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, $"User with username: {username} doesnt exist"));
            }

            var (verified, needsUpgrade) = _passwordHasher.Check(user.Password, password);

            if (!verified)
            {
                context.Status = new Status(StatusCode.InvalidArgument, "Wrong password.");
                return new VerifyUser_Response
                {
                    ErrorMessage = "Wrong password."
                };
            }

            return new VerifyUser_Response
            {
                Verified = true,
                User = _mapper.Map<User>(user)
            };

        }

        public override async Task<ChangeUserInfo_Response> ChangeUserInfo(ChangeUserInfo_Request request, ServerCallContext context)
        {
            using var channel = _grpcChannelBuilder.Build(_servicesConfig.AUTH_SERVICE_ADDRESS);
            var authServiceClient = new AuthService.AuthServiceClient(channel);
            var callOptions = new CallOptions(new Metadata());
            if (context.RequestHeaders.Get("Authorization") != null)
                callOptions.Headers?.Add("Authorization", context.RequestHeaders.Get("Authorization")?.Value);

            var verifyResponse = await authServiceClient.VerifyAsync(new AuthServiceClient.Empty_Request(), callOptions);

            if (!verifyResponse.Verified)
            {
                _logger.LogError(verifyResponse.ErrorMessage);
                var metadata = new Metadata
                {
                    { "UserId", verifyResponse.UserId }
                };
                throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthenticated"), metadata);
            }

            var userToUpdate = await _dbContext.Users.FirstOrDefaultAsync(x => x.Username.Equals(request.Username));
            if (userToUpdate == null)
            {
                _logger.LogInformation($"User with username: {request.Username} doesn't exist.");
                var metadata = new Metadata
                {
                    { "Username", request.Username }
                };
                throw new RpcException(new Status(StatusCode.NotFound, "User doesn't exist."), metadata);
            }

            if (!userToUpdate.Id.Equals(Guid.Parse(verifyResponse.UserId)))
            {
                _logger.LogWarning($"User with id: ${verifyResponse.UserId} tried to update user with id: ${userToUpdate.Id}");
                throw new RpcException(new Status(StatusCode.NotFound, "User doesn't exist."));
            }

            userToUpdate.LivingPlace = request.LivingPlace.Length == 0 ? userToUpdate.LivingPlace : request.LivingPlace;
            userToUpdate.FirstName = request.FirstName.Length == 0 ? userToUpdate.FirstName : request.FirstName;
            userToUpdate.LastName = request.LastName.Length == 0 ? userToUpdate.LastName : request.LastName;

            await _dbContext.SaveChangesAsync();

            return new ChangeUserInfo_Response
            {
                Success = true,
                User = _mapper.Map<UserLessInfo>(userToUpdate)
            };
        }
    }
}

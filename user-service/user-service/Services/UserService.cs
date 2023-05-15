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
        private readonly UserServiceDbContext _dbContext;
        private readonly IMapper _mapper;
        private readonly IValidator<RegisterUser_Request> _validator;
        private readonly IPasswordHasher _passwordHasher;

        public UserService(UserServiceDbContext dbContext,
            IMapper mapper,
            IValidator<RegisterUser_Request> validator,
            IPasswordHasher passwordHasher)
        {
            _dbContext = dbContext;
            _mapper = mapper;
            _validator = validator;
            _passwordHasher = passwordHasher;
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
                errorResponse.Errors.AddRange(validationResult.Errors.Select(x => new user_service.Error
                {
                    PropertyName = x.PropertyName,
                    ErrorMessage = x.ErrorMessage
                }));

                context.Status = new Status(StatusCode.InvalidArgument, "Not all fields are valid.");
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
                        Message = "User successfully registred",
                        User = _mapper.Map<user_service.User>(userEntity)
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

                response.Users.AddRange(users.Select(x => _mapper.Map<user_service.GetAllUsers_Response.Types.User>(x)));

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
            try
            {
                var username = request.Username;
                var password = request.Password;
                var user = await _dbContext.Users
                    .Where(x => x.Username == username)
                    .FirstOrDefaultAsync();

                if (user == null)
                {
                    context.Status = new Status(StatusCode.NotFound, $"User with username: {username} doesnt exist");
                    return new VerifyUser_Response
                    {
                        ErrorMessage = $"User with username: {username} doesnt exist"
                    };
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
            catch (Exception ex)
            {
                if (ex is RpcException rpcException)
                {
                    throw new RpcException(rpcException.Status, rpcException.Trailers);
                }
                throw new RpcException(new Status(StatusCode.Internal, ex.InnerException?.Message ?? ex.Message), Metadata.Empty);
            }

        }
    }
}

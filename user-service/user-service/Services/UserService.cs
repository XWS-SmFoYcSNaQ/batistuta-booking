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
using user_service.domain.Enums;
using System.Threading.Channels;
using Google.Protobuf.WellKnownTypes;

namespace user_service.Services
{
    public class UserService : user_service.UserService.UserServiceBase
    {
        private readonly UserServiceDbContext _dbContext;
        private readonly IMapper _mapper;
        private readonly ILogger<UserService> _logger;
        private readonly IValidator<RegisterUser_Request> _registerRequestValidator;
        private readonly IValidator<ChangePassword_Request> _changePasswordRequestValidator;
        private readonly IPasswordHasher _passwordHasher;
        private readonly GrpcChannelBuilder _grpcChannelBuilder;
        private readonly ServicesConfig _servicesConfig;


        public UserService(UserServiceDbContext dbContext,
            IMapper mapper,
            ILogger<UserService> logger,
            IValidator<RegisterUser_Request> registerRequestValidator,
            IValidator<ChangePassword_Request> changePasswordRequestValidator,
            IPasswordHasher passwordHasher,
            GrpcChannelBuilder grpcChannelBuilder,
            ServicesConfig servicesConfig)
        {
            _dbContext = dbContext;
            _mapper = mapper;
            _logger = logger;
            _registerRequestValidator = registerRequestValidator;
            _changePasswordRequestValidator = changePasswordRequestValidator;
            _passwordHasher = passwordHasher;
            _grpcChannelBuilder = grpcChannelBuilder;
            _servicesConfig = servicesConfig;
        }

        public override async Task<RegisterUser_Response> RegisterUser(RegisterUser_Request request, ServerCallContext context)
        {
            var validationResult = await _registerRequestValidator.ValidateAsync(request);

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

        public override async Task<GetAllUsers_Response> GetAllUsers(Empty_Message request, ServerCallContext context)
        {
            try
            {
                var users = await _dbContext.Users.ToListAsync();
                var response = new GetAllUsers_Response();

                response.Users.AddRange(users.Select(x => _mapper.Map<User>(x)));

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

            var verifyResponse = await authServiceClient.VerifyAsync(new Empty_Request(), new CallOptions().WithHeaders(context.RequestHeaders));

            if (!verifyResponse.Verified)
            {
                _logger.LogError(verifyResponse.ErrorMessage);
                throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthenticated"));
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

        public override async Task<Empty_Message> ChangePassword(ChangePassword_Request request, ServerCallContext context)
        {
            var validatonResult = await _changePasswordRequestValidator.ValidateAsync(request);

            if (!validatonResult.IsValid)
            {
                var metadata = new Metadata();
                validatonResult.Errors.ForEach(x =>
                {
                    metadata.Add(x.PropertyName, x.ErrorMessage);
                });

                throw new RpcException(new Status(StatusCode.InvalidArgument, "Error new password is invalid."), metadata);
            }

            using var channel = _grpcChannelBuilder.Build(_servicesConfig.AUTH_SERVICE_ADDRESS);
            var authServiceClient = new AuthService.AuthServiceClient(channel);

            var verifyResponse = await authServiceClient.VerifyAsync(new Empty_Request(), new CallOptions().WithHeaders(context.RequestHeaders));

            if (!verifyResponse.Verified)
            {
                _logger.LogError(verifyResponse.ErrorMessage);
                throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthenticated"));
            }

            var user = await _dbContext.Users.FirstOrDefaultAsync(x => x.Id == Guid.Parse(verifyResponse.UserId));

            if (user == null)
            {
                _logger.LogError($"Failed changing password, user with id: {verifyResponse.UserId} not found.");
                throw new RpcException(new Status(StatusCode.NotFound, "User doesn't exist"));
            }


            var (verified, needsUpgrade) = _passwordHasher.Check(user.Password, request.CurrentPassword);
            if (!verified)
            {
                _logger.LogError($"Error, wrong password.");
                throw new RpcException(new Status(StatusCode.InvalidArgument, "Wrong password"));
            }

            user.Password = _passwordHasher.Hash(request.NewPassword);
            await _dbContext.SaveChangesAsync();

            return new Empty_Message();
        }

        public override async Task<GetAllHostsWithRatings_Response> GetAllHostsWithRatings(
            Empty_Message request,
            ServerCallContext context)
        {
            var response = new GetAllHostsWithRatings_Response();
            var hosts = await _dbContext.Users.Where(x => x.Role == domain.Enums.UserRole.Host).ToListAsync();

            var channel = _grpcChannelBuilder.Build(_servicesConfig.RATING_SERVICE_ADDRESS);
            var ratingClient = new rating_service.RatingService.RatingServiceClient(channel);
            var hostsRatingsResponse = await ratingClient.GetHostRatingsAsync(new rating_service.Empty(), new CallOptions());
            var hostsRatingsGroupedByHost = hostsRatingsResponse.Data
                .ToLookup(x => x.TargetId);

            foreach (var host in hosts)
            {
                var hostWithRating = new HostWithRating
                {
                    Id = host.Id.ToString(),
                    Email = host.Email,
                    FirstName = host.FirstName,
                    LastName = host.LastName,
                    LivingPlace = host.LivingPlace,
                    Username = host.Username,
                    Role = (UserRole)host.Role,
                    Featured = host.Featured.HasValue ? host.Featured.Value : false
                };

                hostWithRating.Ratings.AddRange(hostsRatingsGroupedByHost[host.Id.ToString()]
                    .Select(x => _mapper.Map<user_service.RatingDTO>(x)));

                response.Hosts.Add(hostWithRating);
            }

            return response;
        }

        public override async Task<Empty_Message> UpdateHostFeatured(UpdateHostFeatured_Request request, ServerCallContext context)
        {
            var hostRating = await _dbContext
                .HostRatings
                .Include(x => x.Host)
                .FirstOrDefaultAsync(x => x.HostId == Guid.Parse(request.HostId));

            if (hostRating == null || hostRating.AverageRating <= 4.7)
            {
                return new Empty_Message();
            }


            var bookingChannel = _grpcChannelBuilder.Build(_servicesConfig.BOOKING_SERVICE_ADDRESS);
            var bookingClient = new booking_service.BookingService.BookingServiceClient(bookingChannel);

            var hostFeaturedResponse = await bookingClient.HostStandOutCheckAsync(new booking_service.EmptyMessage());

            if (!hostFeaturedResponse.Flag)
            {
                _logger.LogInformation(hostFeaturedResponse.Message);
                return new Empty_Message();
            }

            hostRating.Host.Featured = true;
            await _dbContext.SaveChangesAsync();

            return new Empty_Message();
        }

        public override async Task<GetFeaturedHosts_Response> GetFeaturedHosts(Empty_Message request, ServerCallContext context)
        {
            var featuredHosts = await _dbContext.Users.Where(x => x.Role == domain.Enums.UserRole.Host && x.Featured == true).ToListAsync();
            var response = new GetFeaturedHosts_Response();
            response.Hosts.AddRange(featuredHosts.Select(x => _mapper.Map<User>(x)));

            return response;
        }
    }
}

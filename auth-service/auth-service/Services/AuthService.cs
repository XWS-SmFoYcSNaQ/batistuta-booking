﻿using auth_service.Configuration;
using auth_service.Helpers;
using AutoMapper;
using Grpc.Core;
using Grpc.Net.Client;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using UserServiceClient;

namespace auth_service.Services
{
    public class AuthService : auth_service.AuthService.AuthServiceBase
    {
        private readonly ILogger<AuthService> _logger;
        private readonly IMapper _mapper;
        private readonly ServicesConfig _servicesConfig;
        private readonly JwtSettings _jwtSettings;
        private readonly TokenValidationParameters _tokenValidationParameters;
        private readonly GrpcChannelBuilder _grpcChannelBuilder;

        public AuthService(ILogger<AuthService> logger,
            IMapper mapper,
            ServicesConfig servicesConfig,
            JwtSettings jwtSettings,
            TokenValidationParameters tokenValidationParameters,
            GrpcChannelBuilder grpcChannelBuilder)
        {
            _logger = logger;
            _mapper = mapper;
            _servicesConfig = servicesConfig;
            _jwtSettings = jwtSettings;
            _tokenValidationParameters = tokenValidationParameters;
            _grpcChannelBuilder = grpcChannelBuilder;
        }

        public override async Task<Register_Response> Register(Register_Request request, ServerCallContext context)
        {
            using var channel = _grpcChannelBuilder.Build(_servicesConfig.USER_SERVICE_ADDRESS);
            var client = new UserService.UserServiceClient(channel);
            var registerUserRequest = _mapper.Map<RegisterUser_Request>(request);
            var response = await client.RegisterUserAsync(registerUserRequest);

            if (response == null || !response.Success)
            {
                _logger.LogInformation($"Failed to register user: {request.Username}.");
                var registerResponse = new Register_Response()
                {
                    Message = response?.Message
                };

                registerResponse.Errors.AddRange(response?.Errors.Select(_mapper.Map<Error>));
                context.Status = new Status(StatusCode.Unknown, response?.Message ?? "Unknown error occured.");
                return registerResponse;
            }
            _logger.LogInformation($"User: {request.Username} successfully registred.");
            var authenticationResponse = GenerateAuthenticationResponseForUser(response.User);
            return new Register_Response
            {
                Success = authenticationResponse.Success,
                Message = "Registration successfull.",
                Token = authenticationResponse.Token,
                User = _mapper.Map<User>(response.User)
            };
        }

        public override async Task<Authentication_Response> Login(Authentication_Request request, ServerCallContext context)
        {
            using var channel = _grpcChannelBuilder.Build(_servicesConfig.USER_SERVICE_ADDRESS);
            var client = new UserService.UserServiceClient(channel);
            var verifyUserResponse = await client.VerifyUserPasswordAsync(_mapper.Map<VerifyUser_Request>(request));

            if (!verifyUserResponse.Verified || verifyUserResponse.User == null)
            {
                _logger.LogError($"Failed to login, user id: {verifyUserResponse.User}");
                context.Status = new Status(StatusCode.InvalidArgument, verifyUserResponse.ErrorMessage ?? "Error");
                return new Authentication_Response
                {
                    ErrorMessage = verifyUserResponse.ErrorMessage ?? "Error"
                };
            }

            return GenerateAuthenticationResponseForUser(verifyUserResponse.User);
        }

        public override Task<Verify_Response> Verify(Empty_Request request, ServerCallContext context)
        {
            var authorizationHeader = context.RequestHeaders.Get("Authorization")?.Value;
            if (authorizationHeader == null)
            {
                context.Status = new Status(StatusCode.Unauthenticated, "Missing authorization header");
                return Task.FromResult(new Verify_Response
                {
                    ErrorMessage = "Missing authorization header"
                });
            }
            var jwt = authorizationHeader.Length > 7 ? authorizationHeader.Substring(7) : authorizationHeader;

            var claimsPrinciple = VerifyJwt(jwt);

            if (claimsPrinciple == null)
            {
                context.Status = new Status(StatusCode.Unauthenticated, "Invalid token");
                return Task.FromResult(new Verify_Response
                {
                    ErrorMessage = "Invalid token"
                });
            }
            return Task.FromResult(new Verify_Response
            {
                Verified = true,
                UserId = claimsPrinciple.Claims.Single(x => x.Type == "user_id").Value,
                UserRole = Enum.Parse<UserRole>(claimsPrinciple.Claims.Single(x => x.Type == "user_role").Value)
            });
        }

        private Authentication_Response GenerateAuthenticationResponseForUser(UserServiceClient.User newUser)
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var key = Encoding.ASCII.GetBytes(_jwtSettings.Secret!);
            var tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(new[]
                {
                    new Claim(JwtRegisteredClaimNames.Sub, newUser.Username),
                    new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
                    new Claim(JwtRegisteredClaimNames.Email, newUser.Email),
                    new Claim("user_id", newUser.Id.ToString()),
                    new Claim("user_role", newUser.Role.ToString())
                }),
                Expires = DateTime.UtcNow.AddDays(1),
                SigningCredentials = new SigningCredentials(new SymmetricSecurityKey(key), SecurityAlgorithms.HmacSha256Signature)
            };

            var token = tokenHandler.CreateToken(tokenDescriptor);

            return new Authentication_Response
            {
                Success = true,
                Token = tokenHandler.WriteToken(token),
                User = _mapper.Map<User>(newUser)
            };
        }

        private ClaimsPrincipal? VerifyJwt(string jwt)
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            try
            {
                var principle = tokenHandler.ValidateToken(jwt, _tokenValidationParameters, out var validatedToken);
                if (!IsJwtWithValidSecurityAlgorithm(validatedToken))
                {
                    return null;
                }

                return principle;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex.ToString());
                return null;
            }
        }

        private bool IsJwtWithValidSecurityAlgorithm(SecurityToken validatedToken)
        {
            return (validatedToken is JwtSecurityToken jwtSecurityToken) &&
                jwtSecurityToken.Header.Alg.Equals(SecurityAlgorithms.HmacSha256, StringComparison.InvariantCultureIgnoreCase);
        }
    }
}

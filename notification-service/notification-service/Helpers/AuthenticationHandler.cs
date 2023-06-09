﻿using Grpc.Core;
using Microsoft.AspNetCore.Authentication;
using Microsoft.Extensions.Options;
using notification_service.Configuration;
using System.Security.Claims;
using System.Text.Encodings.Web;

namespace notification_service.Helpers
{
    public class AuthenticationHandler : AuthenticationHandler<AuthenticationSchemeOptions>
    {
        private readonly GrpcChannelBuilder _grpcChannelBuilder;
        private readonly IOptions<ServicesConfig> _servicesConfig;

        public AuthenticationHandler(IOptionsMonitor<AuthenticationSchemeOptions> options,
            ILoggerFactory logger,
            UrlEncoder encoder,
            ISystemClock clock,
            GrpcChannelBuilder grpcChannelBuilder,
            IOptions<ServicesConfig> servicesConfig) : base(options, logger, encoder, clock)
        {
            _grpcChannelBuilder = grpcChannelBuilder;
            _servicesConfig = servicesConfig;
        }

        protected override async Task<AuthenticateResult> HandleAuthenticateAsync()
        {
            using var channel = _grpcChannelBuilder.Build(_servicesConfig.Value.AuthServiceAddress);
            var authClient = new AuthServiceClient.AuthService.AuthServiceClient(channel);

            var accessToken = $"Bearer {Context.Request.Query["access_token"].ToString()}";

            if (!Context.Request.Path.ToString().StartsWith("/hubs"))
            {
                accessToken = Context.Request.Headers.Authorization.ToString();
            }

            if (accessToken == null)
            {
                Logger.LogError("Token missing");
                return AuthenticateResult.Fail("Token missing");
            }

            var verifyRequestMetadata = new Metadata
            {
                new Metadata.Entry("Authorization", accessToken)
            };
            var verifyResponse = await authClient.VerifyAsync(new AuthServiceClient.Empty_Request(), new CallOptions().WithHeaders(verifyRequestMetadata));

            if (verifyResponse == null || !verifyResponse.Verified)
            {
                Logger.LogError("Unauthenticated");
                return AuthenticateResult.Fail(verifyResponse?.ErrorMessage ?? "Unauthenticated");
            };

            Logger.LogInformation($"Connected as {verifyResponse.UserId}");
            var userIdClaim = new Claim("userId", verifyResponse.UserId);
            var userRoleClaim = new Claim("userRole", verifyResponse.UserRole.ToString());
            var claimsIdentity = new ClaimsIdentity(new[] { userIdClaim, userRoleClaim }, "MyScheme");
            var claimsPrinciple = new ClaimsPrincipal(claimsIdentity);
            var authenticationTicket = new AuthenticationTicket(claimsPrinciple, "MyScheme");

            return AuthenticateResult.Success(authenticationTicket);
        }
    }
}

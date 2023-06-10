using auth_service.Configuration;
using auth_service.Helpers;
using Grpc.Core;
using Grpc.Net.Client;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using System.Text;

namespace auth_service.Extensions
{
    public static class ConfigurationExetensions
    {
        public static void AddServicesConfig(this WebApplicationBuilder builder)
        {
            var servicesConfig = new ServicesConfig();
            builder.Configuration.Bind("Services", servicesConfig);
            builder.Services.AddSingleton(servicesConfig);
        }

        public static void AddJwtAuthentication(this WebApplicationBuilder builder)
        {
            var jwtSettings = new JwtSettings();
            builder.Configuration.Bind(nameof(jwtSettings), jwtSettings);
            builder.Services.AddSingleton(jwtSettings);

            var tokenValidationParameters = new TokenValidationParameters()
            {
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(jwtSettings.Secret!)),
                ValidateAudience = false,
                ValidateLifetime = true,
                ValidateIssuer = false,
                RequireExpirationTime = false
            };
            builder.Services.AddSingleton(tokenValidationParameters);

            builder.Services.AddAuthentication(opts =>
            {
                opts.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
                opts.DefaultScheme = JwtBearerDefaults.AuthenticationScheme;
                opts.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
            }).AddJwtBearer(opts =>
            {
                opts.SaveToken = false;
                opts.TokenValidationParameters = tokenValidationParameters;
            });
        }

        public static void AddGrpcChannelOptions(this WebApplicationBuilder builder)
        {
            var grpcChannelOptions = new GrpcChannelOptions
            {
                Credentials = ChannelCredentials.Insecure
            };
            builder.Services.AddSingleton(grpcChannelOptions);
        }

        public static void AddHelpers(this WebApplicationBuilder builder)
        {
            builder.Services.AddSingleton<GrpcChannelBuilder>();
        }
    }
}

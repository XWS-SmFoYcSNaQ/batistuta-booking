using Grpc.Net.Client;
using Helpers;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.SignalR;
using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using notification_service.Configuration;
using notification_service.Helpers;
using notification_service.HostedServices;
using notification_service.messaging;
using notification_service.messaging.Interfaces;
using notification_service.Repositories;

namespace notification_service.Extensions
{
    public static class ConfigurationExtensions
    {
        public static void AddNotificationDbSettings(this IServiceCollection services, ConfigurationManager configurationManager)
        {
            var dbConfig = new DbConfiguration();
            configurationManager.Bind("DbConfig", dbConfig);
            var connString = $"mongodb://{dbConfig.User}:{dbConfig.Password}@{dbConfig.Server}:{dbConfig.Port}";
            var notificationDbSettings = new NotificationDbSettings
            {
                ConnectionString = connString,
                DatabaseName = dbConfig.DatabaseName,
                UserNotificationsCollectionName = "UserNotifications",
                UserNotificationOptionsCollectionName = "UserNotificationOptions"
            };

            services.AddSingleton(notificationDbSettings);
        }

        public static void AddServices(this IServiceCollection services)
        {
            services.AddSingleton<INotificationRepository, NotificationRepository>();
            services.AddSingleton<INatsClient, NatsClient>();
            services.AddSingleton<IUserIdProvider, UserIdProvider>();
            services.AddSingleton<GrpcChannelBuilder>();
        }

        public static void AddHostedServices(this IServiceCollection services)
        {
            services.AddHostedService<NotificationService>();
        }

        public static void AddAuth(this IServiceCollection services)
        {
            services.AddAuthentication("MyScheme")
                .AddScheme<AuthenticationSchemeOptions, AuthenticationHandler>("MyScheme", o =>
                {

                });
        }

        public static void AddGrpcChannelOptions(this IServiceCollection services)
        {
            var grpcChannelOptions = new GrpcChannelOptions
            {
                Credentials = Grpc.Core.ChannelCredentials.Insecure
            };

            services.AddSingleton(grpcChannelOptions);
        }

    }
}

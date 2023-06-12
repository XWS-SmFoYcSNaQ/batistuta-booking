using Grpc.Net.Client;
using Microsoft.EntityFrameworkCore;
using user_service.Configuration;
using user_service.data.Db;
using user_service.Helpers;

namespace user_service.Extensions
{
    public static class ConfigurationExtensions
    {
        public static void ApplyMigrations(this WebApplication app)
        {
            using var scope = app.Services.CreateScope();

            var services = scope.ServiceProvider;

            var context = services.GetRequiredService<UserServiceDbContext>();
            if (context.Database.GetPendingMigrations().Any())
            {
                context.Database.Migrate();
            }
        }

        public static void AddDb(this WebApplicationBuilder builder)
        {
            var dbConfig = new DbConfig();
            builder.Configuration.GetSection("MySqlConfig").Bind(dbConfig);
            var connString = $"Server={dbConfig.Server};Port={dbConfig.Port};Database={dbConfig.Database};User={dbConfig.User};Password={dbConfig.Password};";
            builder.Services.AddDbContext<UserServiceDbContext>(optBuilder =>
            {
                optBuilder.UseMySQL(connString);
            });
        }

        public static void AddServicesConfig(this WebApplicationBuilder builder)
        {
            var servicesConfig = new ServicesConfig();
            builder.Configuration.Bind("Services", servicesConfig);
            builder.Services.AddSingleton(servicesConfig);
        }

        public static void AddGrpcChannelOptions(this WebApplicationBuilder builder)
        {
            var grpcChannelOptions = new GrpcChannelOptions
            {
                Credentials = Grpc.Core.ChannelCredentials.Insecure
            };

            builder.Services.AddSingleton(grpcChannelOptions);
        }

        public static void AddHelpers(this WebApplicationBuilder builder)
        {
            builder.Services.AddSingleton<GrpcChannelBuilder>();
        }

    }
}

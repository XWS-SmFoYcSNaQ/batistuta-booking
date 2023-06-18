using notification_service.Configuration;
using notification_service.Repositories;

namespace notification_service.Extensions
{
    public static class ConfigurationExtensions
    {
        public static void AddNotificationDbSettings(this IServiceCollection services, ConfigurationManager configurationManager)
        {
            var dbConfig = new DbConfiguration();
            configurationManager.Bind("DbConfig", dbConfig);
            var connString = $"mongodb://{dbConfig.User}:{dbConfig.Password}@{dbConfig.Server}:${dbConfig.Port}";
            var notificationDbSettings = new NotificationDbSettings
            {
                ConnectionString = connString,
                DatabaseName = dbConfig.DatabaseName,
                NotificationsCollectionName = "Notifications"
            };

            services.AddSingleton(notificationDbSettings);
        }

        public static void AddServices(this IServiceCollection services)
        {
            services.AddSingleton<INotificationRepository, NotificationRepository>();
        }

    }
}

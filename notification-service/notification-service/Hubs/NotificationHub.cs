using Helpers;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using notification_service.Domain;
using notification_service.Models;
using notification_service.Repositories;

namespace notification_service.Hubs
{
    [Authorize]
    public class NotificationHub : Hub
    {
        public static readonly ConnectionMapping<Guid> Connections = new ConnectionMapping<Guid>();
        private readonly ILogger<NotificationHub> _logger;

        public IServiceProvider Services { get; }

        public NotificationHub(
            IServiceProvider services,
            ILogger<NotificationHub> logger)
        {
            Services = services;
            _logger = logger;
        }


        public override async Task OnConnectedAsync()
        {
            try
            {
                var userId = Guid.Parse(Context.User.Claims.FirstOrDefault(x => x.Type == "userId").Value);
                var userRole = Context.User.Claims.FirstOrDefault(x => x.Type == "userRole").Value;

                Connections.Add(userId, Context.ConnectionId);

                using var scope = Services.CreateScope();
                var notificationRepo = scope.ServiceProvider.GetRequiredService<INotificationRepository>();

                var userNotificationOptions = await notificationRepo.GetUserNotificationsOptions(userId);

                if (userNotificationOptions == null)
                {
                    userNotificationOptions = new UserNotificationOptionsEntity(userRole)
                    {
                        UserId = userId
                    };
                    await notificationRepo.CreateUserNotificationsOptions(userNotificationOptions);
                }

                var userNotifications = await notificationRepo.GetUserRecentNotifications(userId, userNotificationOptions);

                if (userNotifications != null && userNotifications.Count > 0)
                {
                    foreach (var connection in Connections.GetConnections(userId))
                    {
                        foreach (var userNotification in userNotifications)
                        {
                            userNotification.ReadAt = DateTime.UtcNow;
                            await Clients.Client(connection)
                                .SendAsync("notification", new NotificationClientMessage(userNotification.Notification));
                        }
                    }

                    await notificationRepo.BulkUpdateNotificationsReadTime(userNotifications.Select(x => x.Id).ToHashSet());
                }
            }
            catch (Exception ex)
            {
                _logger.LogError(ex.ToString());
            }

            await base.OnConnectedAsync();
        }


        public override Task OnDisconnectedAsync(Exception? exception)
        {
            var userId = Guid.Parse(Context.User.Claims.FirstOrDefault(x => x.Type == "userId").Value);

            Connections.Remove(userId, Context.ConnectionId);

            return base.OnDisconnectedAsync(exception);
        }
    }
}

using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.Options;
using NATS.Client;
using notification_service.Configuration;
using notification_service.Domain;
using notification_service.Hubs;
using notification_service.messaging.Interfaces;
using notification_service.Models;
using notification_service.Repositories;
using System.Text.Json;

namespace notification_service.HostedServices
{
    public class NotificationService : BackgroundService
    {
        private readonly ILogger<NotificationService> _logger;
        private readonly INatsClient _natsClient;
        private readonly IOptions<NotificationNatsConfig> _notificationSubjects;
        private readonly INotificationRepository _notificationRepository;
        private readonly string _serviceName = nameof(NotificationService);
        private string _queueGroup => _serviceName;
        public IServiceProvider Services { get; }

        public NotificationService(
            ILogger<NotificationService> logger,
            INatsClient natsClient,
            IOptions<NotificationNatsConfig> notificationSubjects,
            INotificationRepository notificationRepository,
            IServiceProvider services)
        {
            _logger = logger;
            _natsClient = natsClient;
            _notificationSubjects = notificationSubjects;
            _notificationRepository = notificationRepository;
            Services = services;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            _logger.LogInformation($"{_serviceName} is starting");

            _natsClient.SubscribeAsync(_notificationSubjects.Value.NotificationSubject, _queueGroup, GetNotificationsHandler());

            await Task.Delay(Timeout.Infinite, stoppingToken);
        }

        private EventHandler<MsgHandlerEventArgs> GetNotificationsHandler()
        {
            return async (sender, e) =>
            {
                try
                {
                    var notificationMessage = JsonSerializer.Deserialize<NotificationMessage>(e.Message.Data);
                    if (notificationMessage is null)
                    {
                        _logger.LogError("Notification message is null");
                        return;
                    }

                    var userNotification = new UserNotificationEntity(notificationMessage);

                    var userNotificationsOptions = await _notificationRepository.GetUserNotificationsOptions(userNotification.NotifierId);

                    if (userNotificationsOptions == null || !userNotificationsOptions.IsNotificationTypeActivated(userNotification.Notification.Type))
                    {
                        _logger.LogInformation($"User with id: {userNotification.NotifierId} has turned of notifications of type: {userNotification.Notification.Type}");
                        return;
                    }

                    using var scope = Services.CreateScope();
                    var hub = scope.ServiceProvider.GetRequiredService<IHubContext<NotificationHub>>();

                    IEnumerable<string>? userConnections;
                    userConnections = NotificationHub.Connections.GetConnections(userNotification.NotifierId);
                    if (userConnections.Any())
                    {
                        foreach (var connection in userConnections)
                        {
                            await hub.Clients.Client(connection)
                                .SendAsync("Notification", new NotificationClientMessage(userNotification));
                        }
                        userNotification.ReadAt = DateTime.UtcNow;
                    }

                    await _notificationRepository.CreateNotification(userNotification);
                }
                catch (Exception ex) when (ex is JsonException || ex is NotSupportedException)
                {
                    _logger.LogError(ex, "Erorr deserializing notification message.");
                }
                catch (Exception ex)
                {
                    _logger.LogError(ex, "Unknown error occured while processing notification");
                }
            };
        }

        public override Task StopAsync(CancellationToken cancellationToken)
        {
            _logger.LogInformation($"{_serviceName} is stopping");
            return base.StopAsync(cancellationToken);
        }
    }
}

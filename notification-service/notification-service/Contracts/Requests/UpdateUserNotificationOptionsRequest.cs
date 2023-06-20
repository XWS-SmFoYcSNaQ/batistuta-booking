using notification_service.Domain.Enums;

namespace notification_service.Contracts.Requests
{
    public class UpdateUserNotificationOptionsRequest
    {
        public HashSet<NotificationType> ActivedNotifications { get; set; } = new HashSet<NotificationType>();

    }
}

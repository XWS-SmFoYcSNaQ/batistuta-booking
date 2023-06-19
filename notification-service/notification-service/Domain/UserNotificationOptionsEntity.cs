using notification_service.Domain.Enums;

namespace notification_service.Domain
{
    public class UserNotificationOptionsEntity
    {
        public Dictionary<NotificationType, bool> Options { get; set; } = new Dictionary<NotificationType, bool>();
    }
}

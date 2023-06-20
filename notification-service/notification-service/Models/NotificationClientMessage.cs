using notification_service.Domain;
using notification_service.Domain.Enums;

namespace notification_service.Models
{
    public class NotificationClientMessage
    {
        public string Id { get; set; }
        public string Title { get; set; }
        public string Content { get; set; }
        public NotificationType Type { get; set; }
        public DateTime CreatedAt { get; set; }
        public bool New { get; set; }

        public NotificationClientMessage(UserNotificationEntity userNotificationEntity)
        {
            Id = userNotificationEntity.Notification.Id.ToString();
            Title = userNotificationEntity.Notification.Title;
            Content = userNotificationEntity.Notification.Content;
            Type = userNotificationEntity.Notification.Type;
            CreatedAt = userNotificationEntity.Notification.CreatedAt.ToLocalTime();
            New = userNotificationEntity.ReadAt == null;
        }
    }
}

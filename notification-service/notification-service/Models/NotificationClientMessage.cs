using notification_service.Domain;
using notification_service.Domain.Enums;

namespace notification_service.Models
{
    public class NotificationClientMessage
    {
        public string Title { get; set; }
        public string Content { get; set; }
        public NotificationType Type { get; set; }
        public DateTime CreatedAt { get; set; }

        public NotificationClientMessage(NotificationEntity notificationEntity)
        {
            Title = notificationEntity.Title;
            Content = notificationEntity.Content;
            Type = notificationEntity.Type;
            CreatedAt = notificationEntity.CreatedAt;
        }
    }
}

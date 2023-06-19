using notification_service.Domain.Enums;

namespace notification_service.Models
{
    public class NotificationMessage
    {
        public required string Title { get; set; }
        public required string Content { get; set; }
        public required NotificationType Type { get; set; }
        public required Guid NotifierId { get; set; }
        public Guid? ActorId { get; set; }
    }
}

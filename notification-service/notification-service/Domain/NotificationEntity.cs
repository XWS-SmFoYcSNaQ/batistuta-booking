using MongoDB.Bson;
using notification_service.Domain.Enums;

namespace notification_service.Domain
{
    public class NotificationEntity
    {
        public ObjectId Id { get; set; }
        public required string Title { get; set; }
        public required string Content { get; set; }
        public required NotificationType Type { get; set; }
        public required DateTime CreatedAt { get; set; }
    }
}

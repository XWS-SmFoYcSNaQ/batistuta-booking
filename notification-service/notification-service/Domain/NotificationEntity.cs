using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using notification_service.Domain.Enums;
using notification_service.Models;

namespace notification_service.Domain
{
    public class NotificationEntity
    {
        [BsonId]
        public ObjectId Id { get; set; }
        public required string Title { get; set; }
        public required string Content { get; set; }
        public required NotificationType Type { get; set; }
        public required DateTime CreatedAt { get; set; }

    }
}

using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using notification_service.Models;

namespace notification_service.Domain
{
    public class UserNotificationEntity
    {
        [BsonId]
        public ObjectId Id { get; set; }
        public Guid NotifierId { get; set; }
        public Guid? ActorId { get; set; }
        public NotificationEntity Notification { get; set; }
        public DateTime? ReadAt { get; set; }



        public UserNotificationEntity(NotificationMessage notificationMessage)
        {
            Id = ObjectId.GenerateNewId();
            NotifierId = notificationMessage.NotifierId;
            ActorId = notificationMessage.ActorId;
            Notification = new NotificationEntity
            {
                Id = ObjectId.GenerateNewId(),
                Title = notificationMessage.Title,
                Content = notificationMessage.Content,
                Type = notificationMessage.Type,
                CreatedAt = DateTime.UtcNow
            };
        }
    }
}

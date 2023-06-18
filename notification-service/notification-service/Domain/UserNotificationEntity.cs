using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace notification_service.Domain
{
    public class UserNotificationEntity
    {
        public ObjectId Id { get; set; }

        [BsonGuidRepresentation(GuidRepresentation.Standard)]
        public required Guid NotifierId { get; set; }

        [BsonGuidRepresentation(GuidRepresentation.Standard)]
        public Guid? ActorId { get; set; }
        public required NotificationEntity Notification { get; set; }
        public DateTime? ReadAt { get; set; }
    }
}

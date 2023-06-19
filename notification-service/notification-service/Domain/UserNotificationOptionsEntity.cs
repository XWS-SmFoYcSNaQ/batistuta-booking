using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using MongoDB.Bson.Serialization.Options;
using notification_service.Domain.Enums;
using notification_service.Helpers;

namespace notification_service.Domain
{
    public class UserNotificationOptionsEntity
    {
        [BsonId]
        public ObjectId Id { get; set; }
        public required Guid UserId { get; set; }
        public HashSet<NotificationType> ActivatedNotifications { get; set; }

        public UserNotificationOptionsEntity()
        {
            Id = ObjectId.GenerateNewId();
            ActivatedNotifications = new HashSet<NotificationType>();
        }

        public UserNotificationOptionsEntity(string userRole)
        {
            ActivatedNotifications = new HashSet<NotificationType>();

            if (userRole == "Guest")
            {
                ActivatedNotifications.Add(NotificationType.ReservationRequestResponded);
            }
            if (userRole == "Host")
            {
                ActivatedNotifications.Add(NotificationType.ReservationRequestCreated);
                ActivatedNotifications.Add(NotificationType.ReservationCancelled);
                ActivatedNotifications.Add(NotificationType.HostRated);
                ActivatedNotifications.Add(NotificationType.AccommodationRated);
                ActivatedNotifications.Add(NotificationType.HostFeaturedStatusChanged);
            }
        }

        public bool IsNotificationTypeActivated(NotificationType type)
        {
            return ActivatedNotifications.Contains(type);
        }
    }
}

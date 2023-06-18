using MongoDB.Bson;
using notification_service.Domain;

namespace notification_service.Repositories
{
    public interface INotificationRepository
    {
        public Task<List<UserNotificationEntity>> GetUserUnreadNotifications(Guid userId);
        public Task CreateNotification(NotificationEntity notification);
        public Task UpdateNotification(ObjectId notificationId, NotificationEntity updatedNotification);
    }
}

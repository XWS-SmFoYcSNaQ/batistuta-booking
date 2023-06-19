using MongoDB.Bson;
using notification_service.Domain;

namespace notification_service.Repositories
{
    public interface INotificationRepository
    {
        public Task<List<UserNotificationEntity>> GetUserUnreadNotifications(Guid userId);
        public Task CreateNotification(UserNotificationEntity notification);
        public Task UpdateNotification(ObjectId userNotificationId, UserNotificationEntity updatedUserNotification);
        public Task BulkUpdateNotificationsReadTime(HashSet<ObjectId> userNotificationIds);
    }
}

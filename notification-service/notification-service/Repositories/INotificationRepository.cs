using MongoDB.Bson;
using notification_service.Domain;

namespace notification_service.Repositories
{
    public interface INotificationRepository
    {
        public Task<List<UserNotificationEntity>> GetUserUnreadNotifications(Guid userId);
        public Task<List<UserNotificationEntity>> GetUserRecentNotifications(Guid userId, UserNotificationOptionsEntity notificationOptions);
        public Task CreateNotification(UserNotificationEntity notification);
        public Task UpdateNotification(ObjectId userNotificationId, UserNotificationEntity updatedUserNotification);
        public Task BulkUpdateNotificationsReadTime(HashSet<ObjectId> userNotificationIds);

        public Task<UserNotificationOptionsEntity?> GetUserNotificationsOptions(Guid userId);
        public Task CreateUserNotificationsOptions(UserNotificationOptionsEntity userNotificationOptions);
        public Task UpsertUserNotificationsOptions(UserNotificationOptionsEntity userNotificationOptions);
    }
}

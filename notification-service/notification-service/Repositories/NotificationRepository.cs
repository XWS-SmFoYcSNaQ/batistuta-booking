using MongoDB.Bson;
using MongoDB.Driver;
using notification_service.Configuration;
using notification_service.Domain;

namespace notification_service.Repositories
{
    public class NotificationRepository : INotificationRepository
    {
        private readonly IMongoCollection<NotificationEntity> _notificationsCollection;
        private readonly IMongoCollection<UserNotificationEntity> _userNotificationsCollection;

        public NotificationRepository(NotificationDbSettings notificationDbSettings)
        {
            var mongoClient = new MongoClient(notificationDbSettings.ConnectionString);
            var mongoDatabase = mongoClient.GetDatabase(notificationDbSettings.DatabaseName);

            _notificationsCollection = mongoDatabase.GetCollection<NotificationEntity>(notificationDbSettings.NotificationsCollectionName);
            _userNotificationsCollection = mongoDatabase.GetCollection<UserNotificationEntity>(notificationDbSettings.UserNotificationsCollectionName);
        }

        public async Task<List<UserNotificationEntity>> GetUserUnreadNotifications(Guid userId)
        {
            return await _userNotificationsCollection.Find(x => x.NotifierId == userId && x.ReadAt == null).ToListAsync();
        }

        public async Task CreateNotification(NotificationEntity notification)
        {
            await _notificationsCollection.InsertOneAsync(notification);
        }

        public async Task UpdateNotification(ObjectId notificationId, NotificationEntity updatedNotification)
        {
            await _notificationsCollection.ReplaceOneAsync(x => x.Id == notificationId, updatedNotification);
        }
    }
}

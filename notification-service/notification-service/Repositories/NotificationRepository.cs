using MongoDB.Bson;
using MongoDB.Driver;
using notification_service.Configuration;
using notification_service.Domain;

namespace notification_service.Repositories
{
    public class NotificationRepository : INotificationRepository
    {
        private readonly IMongoCollection<UserNotificationEntity> _userNotificationsCollection;
        private readonly IMongoCollection<UserNotificationOptionsEntity> _userNotificationOptionsCollection;
        public NotificationRepository(NotificationDbSettings notificationDbSettings)
        {
            var mongoClient = new MongoClient(notificationDbSettings.ConnectionString);
            var mongoDatabase = mongoClient.GetDatabase(notificationDbSettings.DatabaseName);

            _userNotificationsCollection = mongoDatabase.GetCollection<UserNotificationEntity>(notificationDbSettings.UserNotificationsCollectionName);
            _userNotificationOptionsCollection = mongoDatabase.GetCollection<UserNotificationOptionsEntity>(notificationDbSettings.UserNotificationOptionsCollectionName);
        }

        public async Task<List<UserNotificationEntity>> GetUserUnreadNotifications(Guid userId)
        {
            return await _userNotificationsCollection.Find(x => x.NotifierId == userId && x.ReadAt == null).ToListAsync();
        }

        public async Task CreateNotification(UserNotificationEntity notification)
        {
            await _userNotificationsCollection.InsertOneAsync(notification);
        }

        public async Task UpdateNotification(ObjectId userNotificationId, UserNotificationEntity updatedUserNotification)
        {
            await _userNotificationsCollection.ReplaceOneAsync(x => x.Id == userNotificationId, updatedUserNotification);
        }

        public async Task BulkUpdateNotificationsReadTime(HashSet<ObjectId> userNotificationIds)
        {
            var updateDefiniton = Builders<UserNotificationEntity>.Update.Set(notification => notification.ReadAt, DateTime.UtcNow);
            await _userNotificationsCollection.UpdateManyAsync(x => userNotificationIds.Contains(x.Id), updateDefiniton);
        }
    }
}

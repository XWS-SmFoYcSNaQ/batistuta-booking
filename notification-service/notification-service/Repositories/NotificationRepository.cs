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

        public async Task<List<UserNotificationEntity>> GetUserRecentNotifications(Guid userId, UserNotificationOptionsEntity notificationOptions)
        {
            var limit = 15;

            var usersRecentNotifications = await _userNotificationsCollection
                .Find(x => x.NotifierId == userId
                    && notificationOptions.ActivatedNotifications.Contains(x.Notification.Type)
                    && x.ReadAt == null
                    && DateTime.UtcNow.Subtract(TimeSpan.FromDays(30)) < x.Notification.CreatedAt)
                .SortByDescending(x => x.Notification.CreatedAt)
                .ToListAsync();

            if (usersRecentNotifications.Count < limit)
            {
                var userRecentSeenNotifications = await _userNotificationsCollection
                    .Find(x => x.NotifierId == userId
                        && notificationOptions.ActivatedNotifications.Contains(x.Notification.Type)
                        && DateTime.UtcNow.Subtract(TimeSpan.FromDays(30)) < x.Notification.CreatedAt)
                    .SortByDescending(x => x.Notification.CreatedAt)
                    .Limit(limit - usersRecentNotifications.Count)
                    .ToListAsync();

                usersRecentNotifications = usersRecentNotifications.UnionBy(userRecentSeenNotifications, x => x.Id).ToList();
            }

            return usersRecentNotifications;
        }

        public async Task CreateNotification(UserNotificationEntity notification)
        {
            if (notification.Notification.Id == ObjectId.Empty)
                notification.Notification.Id = ObjectId.GenerateNewId();

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

        public async Task<UserNotificationOptionsEntity?> GetUserNotificationsOptions(Guid userId)
        {
            return await _userNotificationOptionsCollection
                .Find(x => x.UserId == userId)
                .FirstOrDefaultAsync();
        }

        public async Task CreateUserNotificationsOptions(UserNotificationOptionsEntity userNotificationOptions)
        {
            await _userNotificationOptionsCollection.InsertOneAsync(userNotificationOptions);
        }

        public async Task UpsertUserNotificationsOptions(UserNotificationOptionsEntity userNotificationOptions)
        {
            if (userNotificationOptions.Id == ObjectId.Empty)
                userNotificationOptions.Id = ObjectId.GenerateNewId();
            await _userNotificationOptionsCollection.ReplaceOneAsync(x => x.Id == userNotificationOptions.Id, userNotificationOptions, new ReplaceOptions { IsUpsert = true });
        }
    }
}

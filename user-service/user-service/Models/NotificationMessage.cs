namespace user_service.Models
{
    public enum NotificationType
    {
        ReservationRequestCreated = 0,
        ReservationCancelled = 1,
        HostRated = 2,
        AccommodationRated = 3,
        HostFeaturedStatusChanged = 4,
        ReservationRequestResponded = 5
    }

    public class NotificationMessage
    {
        public required string Title { get; set; }
        public required string Content { get; set; }
        public required NotificationType Type { get; set; }
        public required Guid NotifierId { get; set; }
        public Guid? ActorId { get; set; }
    }
}

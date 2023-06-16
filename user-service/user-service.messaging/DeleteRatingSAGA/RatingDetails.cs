namespace user_service.messaging.DeleteRatingSAGA
{
    public class RatingDetails
    {
        public required Guid ID { get; set; }
        public required CreateRatingSAGA.RatingDetails OldValue { get; set; }
    }
}

namespace user_service.messaging.CreateRatingSAGA
{
    public class RatingDetails
    {
        public Guid ID { get; set; }
        public Guid TargetID { get; set; }
        public uint TargetType { get; set; }
        public Guid UserID { get; set; }
        public uint Value { get; set; }
        public RatingDetails? OldValue { get; set; }
    }
}

namespace user_service.messaging.CreateRatingSAGA
{
    public enum CreateRatingReplyType
    {
        HostUpdated = 4,
        HostUpdateFailed = 5,
        UnknownReply = 8
    }

    public class CreateRatingReply
    {
        public required RatingDetails Rating { get; set; }
        public required CreateRatingReplyType Type { get; set; }
    }
}

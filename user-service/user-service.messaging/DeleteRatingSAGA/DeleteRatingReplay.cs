namespace user_service.messaging.DeleteRatingSAGA
{
    public enum DeleteRatingReplayType
    {
        HostUpdated = 4,
        HostUpdateFailed = 5,
        UnknownReplay = 8
    }

    public class DeleteRatingReplay
    {
        public required RatingDetails Rating { get; set; }
        public required DeleteRatingReplayType Type { get; set; }
    }
}

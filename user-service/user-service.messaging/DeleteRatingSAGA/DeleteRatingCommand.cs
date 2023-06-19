namespace user_service.messaging.DeleteRatingSAGA
{
    public enum DeleteRatingCommandType
    {
        UpdateHost = 2,
        RollbackRating = 3,
        ConcludeRatingDeletion = 4,
        UnknownCommand = 5
    }

    public class DeleteRatingCommand
    {
        public required RatingDetails Rating { get; set; }
        public required DeleteRatingCommandType Type { get; set; }
    }
}

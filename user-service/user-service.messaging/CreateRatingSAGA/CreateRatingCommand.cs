namespace user_service.messaging.CreateRatingSAGA
{
    public enum CreateRatingCommandType
    {
        UpdateHost = 2,
        RollbackRating = 3,
        UnknownCommand = 5
    }

    public class CreateRatingCommand
    {
        public required RatingDetails Rating { get; set; }
        public required CreateRatingCommandType Type { get; set; }
    }
}

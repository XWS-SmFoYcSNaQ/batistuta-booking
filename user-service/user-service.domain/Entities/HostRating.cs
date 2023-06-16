namespace user_service.domain.Entities
{
    public class HostRating
    {
        public Guid Id { get; set; }
        public double AverageRating { get; set; }
        public uint NumOfRatings { get; set; }
        public uint TotalRating { get; set; }

        public virtual User? Host { get; set; }
        public required Guid HostId { get; set; }
    }
}

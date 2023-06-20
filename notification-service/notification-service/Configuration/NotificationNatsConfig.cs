namespace notification_service.Configuration
{
    public class NotificationNatsConfig
    {
        public string NotificationSubject { get; set; } = string.Empty;
        public string QueueGroup { get; set; } = string.Empty;
    }
}

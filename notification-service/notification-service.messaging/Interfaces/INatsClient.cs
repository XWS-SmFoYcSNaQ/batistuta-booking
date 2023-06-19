using NATS.Client;

namespace notification_service.messaging.Interfaces
{
    public interface INatsClient
    {
        public IAsyncSubscription SubscribeAsync(string topic, string queueGroup, EventHandler<MsgHandlerEventArgs> handler);
    }
}

using NATS.Client;

namespace user_service.messaging.Interfaces
{
    public interface INatsClient : IDisposable
    {
        public IAsyncSubscription SubscribeAsync(string topic, EventHandler<MsgHandlerEventArgs> handler);
        public void Publish(string topic, string message);
    }
}

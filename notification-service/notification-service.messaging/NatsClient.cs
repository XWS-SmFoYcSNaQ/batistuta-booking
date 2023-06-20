using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using NATS.Client;
using notification_service.messaging.Configuration;
using notification_service.messaging.Interfaces;

namespace notification_service.messaging
{
    public class NatsClient : INatsClient
    {
        private readonly IConnection? _connection;
        private readonly IOptions<NatsConfig> _natsConfig;
        private readonly ILogger<NatsClient> _logger;

        public NatsClient(
            IOptions<NatsConfig> natsConfig,
            ILogger<NatsClient> logger
            )
        {
            _natsConfig = natsConfig;
            _logger = logger;
            var connectionFactory = new ConnectionFactory();
            var url = $"nats://${_natsConfig.Value.User}:{_natsConfig.Value.Pass}@{_natsConfig.Value.Host}:{_natsConfig.Value.Port}";
            try
            {
                _connection = connectionFactory.CreateConnection(url);
                _logger.LogInformation("Connected to nats server successfully");
            }
            catch (NATSNoServersException ex)
            {
                _logger.LogError(ex, "Error connecting to nats, no server");
            }
            catch (NATSConnectionException ex)
            {
                _logger.LogError(ex, "Error connecting");
            }
        }

        public IAsyncSubscription SubscribeAsync(string topic, string queueGroup, EventHandler<MsgHandlerEventArgs> handler)
        {
            return _connection.SubscribeAsync(topic, queueGroup, handler);
        }
    }
}

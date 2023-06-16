
using Microsoft.Extensions.Logging;
using NATS.Client;
using System.Text;
using user_service.messaging.Configuration;
using user_service.messaging.Interfaces;

namespace user_service.messaging
{
    public class NatsClient : INatsClient
    {
        private readonly IConnection? _connection;
        private readonly NatsConfiguration _natsConfiguration;
        private readonly ILogger<NatsClient> _logger;

        public NatsClient(NatsConfiguration natsConfiguration, ILogger<NatsClient> logger)
        {
            _natsConfiguration = natsConfiguration;
            _logger = logger;
            var connectionFactory = new ConnectionFactory();
            var url = $"nats://${_natsConfiguration.USER}:{_natsConfiguration.PASSWORD}@{_natsConfiguration.HOST}:{_natsConfiguration.PORT}";
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

        public IAsyncSubscription SubscribeAsync(string subject, EventHandler<MsgHandlerEventArgs> handler)
        {
            return _connection?.SubscribeAsync(subject, handler);
        }

        public void Publish(string subject, string message)
        {
            _connection?.Publish(subject, Encoding.UTF8.GetBytes(message));
        }

        public void Dispose()
        {
            _connection?.Drain();
            _connection?.Close();
        }
    }
}

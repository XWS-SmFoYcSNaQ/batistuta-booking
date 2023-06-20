using Microsoft.EntityFrameworkCore;
using System.Text.Json;
using user_service.Configuration;
using user_service.data.Db;
using user_service.Helpers;
using user_service.messaging.Interfaces;
using user_service.Models;

namespace user_service.Services
{
    public class HostFeaturedUpdater
    {
        private readonly UserServiceDbContext _dbContext;
        private readonly ILogger<HostFeaturedUpdater> _logger;
        private readonly GrpcChannelBuilder _grpcChannelBuilder;
        private readonly ServicesConfig _servicesConfig;
        private readonly INatsClient _natsClient;

        public HostFeaturedUpdater(
            UserServiceDbContext userServiceDbContext,
            ILogger<HostFeaturedUpdater> logger,
            GrpcChannelBuilder grpcChannelBuilder,
            ServicesConfig servicesConfig,
            INatsClient natsClient)
        {
            _dbContext = userServiceDbContext;
            _logger = logger;
            _grpcChannelBuilder = grpcChannelBuilder;
            _servicesConfig = servicesConfig;
            _natsClient = natsClient;
        }



        public async Task UpdateFeatured(Guid hostId)
        {
            var hostRating = await _dbContext
                .HostRatings
                .Include(x => x.Host)
                .FirstOrDefaultAsync(x => x.HostId == hostId);

            if (hostRating == null || hostRating.AverageRating <= 4.7)
            {
                var host = await _dbContext.Users.FirstOrDefaultAsync(x => x.Id == hostId && x.Role == domain.Enums.UserRole.Host);
                if (host != null)
                {
                    host.Featured = false;
                    var rows = await _dbContext.SaveChangesAsync();
                    if (rows > 0)
                    {
                        var notification = new NotificationMessage
                        {
                            Content = "Featured status changed",
                            Title = "Status changed",
                            NotifierId = hostId,
                            Type = NotificationType.HostFeaturedStatusChanged
                        };
                        _natsClient.Publish("notification", JsonSerializer.Serialize(notification));
                    }
                }
                _logger.LogError($"Host with id: {hostId} doesn't have rating greater then 4.7");
                return;
            }


            var bookingChannel = _grpcChannelBuilder.Build(_servicesConfig.BOOKING_SERVICE_ADDRESS);
            var bookingClient = new booking_service.BookingService.BookingServiceClient(bookingChannel);

            var hostFeaturedResponse = await bookingClient.HostStandOutCheckAsync(new booking_service.EmptyMessage());

            if (!hostFeaturedResponse.Flag)
            {
                _logger.LogInformation(hostFeaturedResponse.Message);
                hostRating.Host.Featured = false;
                await _dbContext.SaveChangesAsync();
                return;
            }

            hostRating.Host.Featured = true;
            await _dbContext.SaveChangesAsync();

        }
    }
}

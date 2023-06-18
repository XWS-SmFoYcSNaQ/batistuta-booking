using Microsoft.EntityFrameworkCore;
using NATS.Client;
using System.Text.Json;
using user_service.Configuration;
using user_service.data.Db;
using user_service.messaging.DeleteRatingSAGA;
using user_service.messaging.Interfaces;

namespace user_service.BackgroundServices
{
    public class DeleteRatingService : BackgroundService
    {
        private readonly ILogger<DeleteRatingService> _logger;
        private readonly INatsClient _natsClient;
        private readonly DeleteRatingSubjectsConfig _deleteRatingSubjectsConfig;
        private static readonly string ServiceName = nameof(DeleteRatingService);

        private string DeleteRatingReplaySubject => _deleteRatingSubjectsConfig.DELETE_RATING_REPLY_SUBJECT;
        private string DeleteRatingCommandSubject => _deleteRatingSubjectsConfig.DELETE_RATING_COMMAND_SUBJECT;
        public IServiceProvider Services { get; }

        public DeleteRatingService(
            ILogger<DeleteRatingService> logger,
            IServiceProvider serviceProvider,
            INatsClient natsClient,
            DeleteRatingSubjectsConfig deleteRatingSubjectsConfig)
        {
            _logger = logger;
            _natsClient = natsClient;
            _deleteRatingSubjectsConfig = deleteRatingSubjectsConfig;

            Services = serviceProvider;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            _logger.LogInformation($"{ServiceName} is starting.");

            _natsClient.SubscribeAsync(DeleteRatingCommandSubject, DeleteRatingHandler());

            await Task.Delay(Timeout.Infinite, stoppingToken);
        }

        private EventHandler<MsgHandlerEventArgs> DeleteRatingHandler()
        {
            return async (senders, args) =>
            {
                try
                {
                    var deleteRatingCommand = JsonSerializer.Deserialize<DeleteRatingCommand>(args.Message.Data);
                    if (deleteRatingCommand == null)
                    {
                        var deleteRatingReplay = new DeleteRatingReplay
                        {
                            Rating = new RatingDetails { ID = Guid.Empty, OldValue = new messaging.CreateRatingSAGA.RatingDetails() },
                            Type = DeleteRatingReplayType.UnknownReplay
                        };
                        _natsClient.Publish(DeleteRatingReplaySubject, JsonSerializer.Serialize(deleteRatingReplay));
                        return;
                    }

                    switch (deleteRatingCommand.Type)
                    {
                        case DeleteRatingCommandType.UpdateHost:
                            await UpdateHost(deleteRatingCommand);
                            break;
                        case DeleteRatingCommandType.RollbackRating:
                            await RollbackRating(deleteRatingCommand);
                            break;
                        default:
                            break;
                    }
                }
                catch (Exception ex) when (ex is JsonException || ex is NotSupportedException)
                {
                    _logger.LogError(ex, "Error while parsing json");
                    var deleteRatingReplay = new DeleteRatingReplay
                    {
                        Rating = new RatingDetails { ID = Guid.Empty, OldValue = new messaging.CreateRatingSAGA.RatingDetails() },
                        Type = DeleteRatingReplayType.UnknownReplay
                    };
                    _natsClient.Publish(DeleteRatingReplaySubject, JsonSerializer.Serialize(deleteRatingReplay));
                }
            };
        }

        private async Task UpdateHost(DeleteRatingCommand deleteRatingCommand)
        {
            using var serviceScope = Services.CreateScope();
            var dbContext = serviceScope.ServiceProvider.GetRequiredService<UserServiceDbContext>();

            var host = await dbContext.Users.FirstOrDefaultAsync(x => x.Id == deleteRatingCommand.Rating.OldValue.TargetID && x.Role == domain.Enums.UserRole.Host);

            if (host == null)
            {
                _logger.LogError($"Host update failed, host with id: {deleteRatingCommand.Rating.OldValue.TargetID} doesn't exist.");
                var deleteRatingReplay = new DeleteRatingReplay
                {
                    Rating = deleteRatingCommand.Rating,
                    Type = DeleteRatingReplayType.HostUpdateFailed
                };
                _natsClient.Publish(DeleteRatingReplaySubject, JsonSerializer.Serialize(deleteRatingReplay));
                return;
            }

            var hostRating = await dbContext.HostRatings.FirstOrDefaultAsync(x => x.HostId == deleteRatingCommand.Rating.OldValue.TargetID);

            if (hostRating == null || hostRating.TotalRating == 0)
            {
                _logger.LogError($"Host with id: ${host.Id} doesn't have any ratings so delete command can't be executed.");
                var deleteRatingReplay = new DeleteRatingReplay
                {
                    Rating = deleteRatingCommand.Rating,
                    Type = DeleteRatingReplayType.HostUpdateFailed
                };
                _natsClient.Publish(DeleteRatingReplaySubject, JsonSerializer.Serialize(deleteRatingReplay));
                return;
            }

            hostRating.NumOfRatings--;
            hostRating.TotalRating -= deleteRatingCommand.Rating.OldValue.Value;
            hostRating.AverageRating = hostRating.NumOfRatings == 0 ? 0 : hostRating.TotalRating / hostRating.NumOfRatings;

            await dbContext.SaveChangesAsync();

            var replay = new DeleteRatingReplay
            {
                Rating = deleteRatingCommand.Rating,
                Type = DeleteRatingReplayType.HostUpdated
            };
            _natsClient.Publish(DeleteRatingReplaySubject, JsonSerializer.Serialize(replay));
        }

        private async Task RollbackRating(DeleteRatingCommand deleteRatingCommand)
        {
            using var serviceScope = Services.CreateScope();
            var dbContext = serviceScope.ServiceProvider.GetRequiredService<UserServiceDbContext>();

            var hostRating = await dbContext.HostRatings.FirstOrDefaultAsync(x => x.HostId == deleteRatingCommand.Rating.OldValue.TargetID);

            if (hostRating == null)
            {
                _logger.LogInformation($"Host with id: {deleteRatingCommand.Rating.OldValue.TargetID} didn't have rating.");
                return;
            }

            hostRating.NumOfRatings++;
            hostRating.TotalRating += deleteRatingCommand.Rating.OldValue.Value;
            hostRating.AverageRating = hostRating.TotalRating / hostRating.NumOfRatings;

            await dbContext.SaveChangesAsync();

            _logger.LogInformation("Host rating delete roledback.");
        }


        public override Task StopAsync(CancellationToken cancellationToken)
        {
            _logger.LogInformation($"{ServiceName} is stopping.");
            return base.StartAsync(cancellationToken);
        }
    }
}

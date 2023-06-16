using Microsoft.EntityFrameworkCore;
using NATS.Client;
using System.Text.Json;
using user_service.Configuration;
using user_service.data.Db;
using user_service.domain.Entities;
using user_service.messaging.CreateRatingSAGA;
using user_service.messaging.Interfaces;

namespace user_service.BackgroundServices
{
    public class CreateRatingService : BackgroundService
    {
        private readonly INatsClient _natsClient;
        private readonly CreateRatingSubjectsConfig _subjects;
        private readonly ILogger<CreateRatingService> _logger;
        private readonly string ServiceName = nameof(CreateRatingService);

        private string CreateRatingReplaySubject => _subjects.CREATE_RATING_REPLAY_SUBJECT;
        private string CreateRatingCommandSubject => _subjects.CREATE_RATING_COMMAND_SUBJECT;
        public IServiceProvider Services { get; }

        public CreateRatingService(
            INatsClient natsClient,
            CreateRatingSubjectsConfig createRatingSubjectsConfig,
            ILogger<CreateRatingService> logger,
            IServiceProvider serviceProvider)
        {
            _natsClient = natsClient;
            _subjects = createRatingSubjectsConfig;
            _logger = logger;
            Services = serviceProvider;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            _logger.LogInformation($"{ServiceName} is starting.");

            await BackgroundProcessing(stoppingToken);
        }

        public override async Task StopAsync(CancellationToken cancellationToken)
        {
            _logger.LogInformation($"{ServiceName} is stopping.");

            await base.StopAsync(cancellationToken);
        }

        private async Task BackgroundProcessing(CancellationToken stoppingToken)
        {
            _natsClient.SubscribeAsync(CreateRatingCommandSubject, CreateRatingHandler());

            await Task.Delay(Timeout.Infinite, stoppingToken);

        }

        private EventHandler<MsgHandlerEventArgs> CreateRatingHandler()
        {
            return async (sender, args) =>
            {
                try
                {
                    var createRatingCommand = JsonSerializer.Deserialize<CreateRatingCommand>(args.Message.Data);
                    if (createRatingCommand == null)
                    {
                        var createRatingReplay = new CreateRatingReply
                        {
                            Rating = new RatingDetails(),
                            Type = CreateRatingReplyType.UnknownReply
                        };
                        _natsClient.Publish(CreateRatingReplaySubject, JsonSerializer.Serialize(createRatingReplay));
                        return;
                    }

                    switch (createRatingCommand.Type)
                    {
                        case CreateRatingCommandType.UpdateHost:
                            await UpdateHost(createRatingCommand);
                            break;
                        case CreateRatingCommandType.RollbackRating:
                            await RollbackRating(createRatingCommand);
                            break;
                        default:
                            break;
                    }

                }
                catch (Exception ex) when (ex is JsonException || ex is NotSupportedException)
                {
                    _logger.LogError(ex, "Error parsing json");
                    var createRatingReplay = new CreateRatingReply
                    {
                        Rating = new RatingDetails(),
                        Type = CreateRatingReplyType.UnknownReply
                    };
                    _natsClient.Publish(CreateRatingReplaySubject, JsonSerializer.Serialize(createRatingReplay));
                }
            };
        }

        private async Task RollbackRating(CreateRatingCommand createRatingCommand)
        {
            using var serviceScope = Services.CreateScope();
            var dbContext = serviceScope.ServiceProvider.GetService<UserServiceDbContext>();

            var hostRating = await dbContext.HostRatings.FirstOrDefaultAsync(x => x.HostId == createRatingCommand.Rating.TargetID);
            if (hostRating == null)
            {
                _logger.LogInformation($"Host with id: {createRatingCommand.Rating.TargetID} didn't have rating");
                return;
            }

            if (createRatingCommand.Rating.OldValue != null)
            {
                hostRating.TotalRating = hostRating.TotalRating - createRatingCommand.Rating.Value + createRatingCommand.Rating.OldValue.Value;
            }
            else
            {
                hostRating.NumOfRatings--;
                hostRating.TotalRating = hostRating.TotalRating - createRatingCommand.Rating.Value;
            }
            hostRating.AverageRating = Math.Round((double)hostRating.TotalRating / hostRating.NumOfRatings, 2);
            await dbContext.SaveChangesAsync();

            _logger.LogInformation("Host rating create/update roledback.");
        }

        private async Task UpdateHost(CreateRatingCommand createRatingCommand)
        {
            try
            {
                using var serviceScope = Services.CreateScope();
                var dbContext = serviceScope.ServiceProvider.GetService<UserServiceDbContext>();

                var host = await dbContext.Users.FirstOrDefaultAsync(x =>
                x.Id == createRatingCommand.Rating.TargetID &&
                x.Role == domain.Enums.UserRole.Host);

                if (host == null)
                {
                    _logger.LogError($"Host update failed, host with id: {createRatingCommand.Rating.TargetID} doesn't exist.");
                    var createRatingReplay = new CreateRatingReply
                    {
                        Rating = createRatingCommand.Rating,
                        Type = CreateRatingReplyType.HostUpdateFailed
                    };
                    _natsClient.Publish(CreateRatingReplaySubject, JsonSerializer.Serialize(createRatingReplay));
                    return;
                }

                var hostRating = await dbContext.HostRatings.FirstOrDefaultAsync(x => x.HostId == createRatingCommand.Rating.TargetID);

                if (hostRating == null)
                {
                    var newHostRating = new HostRating
                    {
                        HostId = createRatingCommand.Rating.TargetID,
                        AverageRating = createRatingCommand.Rating.Value,
                        NumOfRatings = 1,
                        TotalRating = createRatingCommand.Rating.Value
                    };
                    dbContext.HostRatings.Add(newHostRating);
                    await dbContext.SaveChangesAsync();
                }
                else
                {
                    if (createRatingCommand.Rating.OldValue != null)
                    {
                        hostRating.TotalRating = hostRating.TotalRating - createRatingCommand.Rating.OldValue.Value + createRatingCommand.Rating.Value;
                    }
                    else
                    {
                        hostRating.NumOfRatings++;
                        hostRating.TotalRating = hostRating.TotalRating + createRatingCommand.Rating.Value;
                    }

                    hostRating.AverageRating = Math.Round((double)hostRating.TotalRating / hostRating.NumOfRatings, 2);
                    await dbContext.SaveChangesAsync();
                }
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Error updating host");
                var createRatingReplay = new CreateRatingReply
                {
                    Rating = createRatingCommand.Rating,
                    Type = CreateRatingReplyType.HostUpdateFailed
                };
                _natsClient.Publish(CreateRatingReplaySubject, JsonSerializer.Serialize(createRatingReplay));
                return;
            }

            var replay = new CreateRatingReply
            {
                Rating = createRatingCommand.Rating,
                Type = CreateRatingReplyType.HostUpdated
            };
            _natsClient.Publish(CreateRatingReplaySubject, JsonSerializer.Serialize(replay));
        }

    }
}

using Microsoft.Extensions.Options;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Bson.Serialization;
using MongoDB.Bson;
using notification_service.Configuration;
using notification_service.Extensions;
using notification_service.Hubs;
using notification_service.messaging.Configuration;
using notification_service.Domain;
using notification_service.Repositories;
using notification_service.Helpers;
using Grpc.Core;
using Microsoft.AspNetCore.Mvc;
using notification_service.Contracts.Requests;
using System.Security.Claims;
using notification_service.Domain.Enums;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Bson congifuration
BsonSerializer.RegisterSerializer(new GuidSerializer(GuidRepresentation.Standard));
BsonDefaults.GuidRepresentationMode = GuidRepresentationMode.V3;
BsonClassMap.RegisterClassMap<NotificationEntity>(initializer =>
{
    initializer.AutoMap();
    initializer.GetMemberMap(x => x.CreatedAt).SetDefaultValue(DateTime.UtcNow);
});

// Add services to the container.
builder.Services.AddCors(options =>
{
    options.AddDefaultPolicy(builder =>
    {
        builder.WithOrigins("http://localhost:5500", "http://localhost:5272", "http://localhost:3000")
               .AllowAnyHeader()
               .AllowAnyMethod()
               .AllowCredentials();
    });
});

builder.Services.AddGrpc();
builder.Services.AddNotificationDbSettings(builder.Configuration);
builder.Services.Configure<NatsConfig>(builder.Configuration.GetSection("NatsConfiguration"));
builder.Services.Configure<NotificationNatsConfig>(builder.Configuration.GetSection("NotificationNatsConfig"));
builder.Services.AddGrpcChannelOptions();
builder.Services.Configure<ServicesConfig>(builder.Configuration.GetSection("Services"));
builder.Services.AddAuth();
builder.Services.AddSignalR(hubOptions =>
{
    hubOptions.ClientTimeoutInterval = TimeSpan.FromHours(1);
    hubOptions.KeepAliveInterval = TimeSpan.FromMinutes(1);
});
builder.Services.AddServices();
builder.Services.AddHostedServices();

var app = builder.Build();

// Configure the HTTP request pipeline.

app.UseAuthentication();
app.UseAuthorization();

app.UseCors();

app.Urls.Add($"http://localhost:{app.Configuration["NOTIFICATION_SERVICE_ADDRESS"]}" ?? "http://localhost:12009");

app.MapHub<NotificationHub>("/hubs/notification");
app.MapPut("/notificationOptions", async (ClaimsPrincipal user,
    UpdateUserNotificationOptionsRequest updateRequest,
    INotificationRepository notificationRepository,
    GrpcChannelBuilder grpcChannelBuilder,
    IOptions<ServicesConfig> servicesConfig,
    ILoggerFactory loggerFactory) =>
{
    var logger = loggerFactory.CreateLogger("UpdateNotificationOptions");
    try
    {
        var userId = user.Claims.FirstOrDefault(x => x.Type == "userId").Value;
        var userRole = user.Claims.FirstOrDefault(x => x.Type == "userRole").Value;

        var userNotificationOptions = await notificationRepository.GetUserNotificationsOptions(Guid.Parse(userId));
        if (userNotificationOptions == null)
        {
            userNotificationOptions = new UserNotificationOptionsEntity
            {
                UserId = Guid.Parse(userId)
            };
        };

        if (userRole.Equals("Guest"))
        {
            if (updateRequest.ActivedNotifications.Contains(NotificationType.ReservationRequestResponded))
                userNotificationOptions.ActivatedNotifications.Add(NotificationType.ReservationRequestResponded);
            else
                userNotificationOptions.ActivatedNotifications.Clear();
        }
        else
        {
            userNotificationOptions.ActivatedNotifications = updateRequest.ActivedNotifications;
            if (userNotificationOptions.ActivatedNotifications.Contains(NotificationType.ReservationRequestResponded))
                userNotificationOptions.ActivatedNotifications.Remove(NotificationType.ReservationRequestResponded);
        }

        await notificationRepository.UpsertUserNotificationsOptions(userNotificationOptions);

        return Results.NoContent();
    }
    catch (Exception ex)
    {
        logger.LogError(ex.ToString());
        return Results.Problem("Error updating notifications options.");
    }


});

app.Run();

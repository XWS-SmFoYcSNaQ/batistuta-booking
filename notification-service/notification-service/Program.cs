using Microsoft.Extensions.Options;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Bson.Serialization;
using MongoDB.Bson;
using notification_service.Configuration;
using notification_service.Extensions;
using notification_service.Hubs;
using notification_service.messaging.Configuration;
using notification_service.Domain;

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

app.Run();

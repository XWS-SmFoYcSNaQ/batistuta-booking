using FluentValidation.AspNetCore;
using user_service.Extensions;
using user_service.Interceptors;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc(opts =>
{
    opts.EnableDetailedErrors = true;
    opts.Interceptors.Add<ExceptionInterceptor>();
});
builder.AddDb();
builder.AddServicesConfig();
builder.AddGrpcChannelOptions();
builder.AddNatsConfig();
builder.AddCreateRatingSubjectsConfig();
builder.AddDeleteRatingSubjectsConfig();
builder.AddHelpers();
builder.Services.AddAutoMapper(typeof(Program));
builder.Services.AddFluentValidationAutoValidation();
builder.Services.AddFluentValidationClientsideAdapters();
builder.AddServices();
builder.AddHostedServices();

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<user_service.Services.UserService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Urls.Add($"http://localhost:{app.Configuration["USER_SERVICE_ADDRESS"]}");
app.ApplyMigrations();
app.Run();



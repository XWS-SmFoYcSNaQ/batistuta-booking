using auth_service.Extensions;
using auth_service.Services;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc(opts =>
{
    opts.EnableDetailedErrors = true;
});
builder.Services.AddAutoMapper(typeof(Program));
builder.AddServicesConfig();
builder.AddJwtAuthentication();

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<AuthService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.UseAuthentication();

app.Run();

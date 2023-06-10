using FluentValidation;
using FluentValidation.AspNetCore;
using Microsoft.EntityFrameworkCore;
using user_service;
using user_service.data.Db;
using user_service.Extensions;
using user_service.Interceptors;
using user_service.Interfaces;
using user_service.Services;
using user_service.Validators;

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
builder.Services.AddAutoMapper(typeof(Program));
// builder.Services.AddMvc().AddFluentValidation(mvcCongif => mvcCongif.RegisterValidatorsFromAssemblyContaining<Program>());
builder.Services.AddFluentValidationAutoValidation();
builder.Services.AddFluentValidationClientsideAdapters();
builder.Services.AddScoped<IValidator<RegisterUser_Request>, RegisterUserRequestValidator>();
builder.Services.AddScoped<IPasswordHasher, PasswordHasher>();

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<user_service.Services.UserService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Urls.Add($"http://localhost:{app.Configuration["USER_SERVICE_ADDRESS"]}");
app.ApplyMigrations();
app.Run();

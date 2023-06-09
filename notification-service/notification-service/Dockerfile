#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["notification-service/notification-service.csproj", "notification-service/"]
COPY ["notification-service.messaging/notification-service.messaging.csproj", "notification-service.messaging/"]
RUN dotnet restore "notification-service/notification-service.csproj"
COPY . .
WORKDIR "/src/notification-service"
RUN dotnet build "notification-service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "notification-service.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "notification-service.dll"]
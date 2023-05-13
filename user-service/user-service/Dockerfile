#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["user-service/user-service.csproj", "user-service/"]
COPY ["user-service.data/user-service.data.csproj", "user-service.data/"]
COPY ["user-service.domain/user-service.domain.csproj", "user-service.domain/"]
RUN dotnet restore "user-service/user-service.csproj"
COPY . .
WORKDIR "/src/user-service"
RUN dotnet build "user-service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "user-service.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "user-service.dll"]
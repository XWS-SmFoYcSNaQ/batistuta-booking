using Microsoft.AspNetCore.SignalR;

namespace Helpers
{
    public class UserIdProvider : IUserIdProvider
    {
        public string? GetUserId(HubConnectionContext connection)
        {
            return connection.User.Claims.FirstOrDefault(x => x.Type == "userId")?.Value;
        }
    }
}

using AutoMapper;

namespace auth_service.MapperProfiles
{
    public class UserProfile : Profile
    {
        public UserProfile()
        {
            CreateMap<UserServiceClient.User, auth_service.User>();
        }
    }
}

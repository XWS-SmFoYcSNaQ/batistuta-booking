using AutoMapper;
using user_service.Models;

namespace user_service.MappingProfiles
{
    public class UserProfile : Profile
    {
        public UserProfile()
        {
            CreateMap<RegisterUser_Request, User>();
            CreateMap<user_service.Models.User, user_service.domain.Entities.User>();
            CreateMap<user_service.domain.Entities.User, user_service.GetAllUsers_Response.Types.User>();
        }
    }
}

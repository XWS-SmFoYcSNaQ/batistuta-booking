using AutoMapper;
using user_service.Models;

namespace user_service.MappingProfiles
{
    public class UserProfile : Profile
    {
        public UserProfile()
        {
            CreateMap<RegisterUser_Request, User>();
            CreateMap<Models.User, domain.Entities.User>();
            CreateMap<domain.Entities.User, UserLessInfo>();
            CreateMap<RegisterUser_Request, User>();
            CreateMap<RegisterUser_Request, Models.User>();
            CreateMap<domain.Entities.User, User>();
        }
    }
}

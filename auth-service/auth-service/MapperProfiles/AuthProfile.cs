using AutoMapper;
using UserServiceClient;

namespace auth_service.MapperProfiles
{
    public class AuthProfile : Profile
    {
        public AuthProfile()
        {
            CreateMap<Register_Request, RegisterUser_Request>();
            CreateMap<Authentication_Request, VerifyUser_Request>();
        }
    }
}

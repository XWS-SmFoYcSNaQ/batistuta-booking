using AutoMapper;

namespace auth_service.MapperProfiles
{
    public class ErrorProfile : Profile
    {
        public ErrorProfile()
        {
            CreateMap<UserServiceClient.Error, auth_service.Error>();
        }
    }
}

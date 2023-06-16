using AutoMapper;

namespace user_service.MappingProfiles
{
    public class RatingProfile : Profile
    {
        public RatingProfile()
        {
            CreateMap<rating_service.RatingDTO, user_service.RatingDTO>();
        }
    }
}

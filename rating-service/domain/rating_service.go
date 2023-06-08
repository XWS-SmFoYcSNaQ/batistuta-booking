package domain

type RatingService struct {
	repository RatingRepository
}

func NewRatingService(repository *RatingRepository) *RatingService {
	return &RatingService{
		repository: *repository,
	}
}

func (service *RatingService) GetAll() (*[]Rating, error) {
	return service.repository.GetAll()
}

package domain

import (
	"github.com/google/uuid"
)

type RatingService struct {
	repository   RatingRepository
	orchestrator *CreateRatingOrchestrator
}

func NewRatingService(repository *RatingRepository, orchestrator *CreateRatingOrchestrator) *RatingService {
	return &RatingService{
		repository:   *repository,
		orchestrator: orchestrator,
	}
}

func (service *RatingService) GetAll() (*[]Rating, error) {
	return service.repository.GetAll()
}

func (service *RatingService) CreateRating(rating *Rating) error {
	(*rating).ID = uuid.New()
	err := service.repository.Insert(rating)
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(rating)
	if err != nil {
		return err
	}
	return nil
}

func (service *RatingService) Delete(rating *Rating) error {
	return service.repository.Delete(rating)
}

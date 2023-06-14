package domain

import "github.com/google/uuid"

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

func (service *RatingService) Insert(rating *Rating) error {
	return service.repository.Insert(rating)
}

func (service *RatingService) Update(rating *Rating) error {
	return service.repository.Update(rating)
}

func (service *RatingService) CreateRating(rating *Rating) error {
	var oldRating *Rating = nil
	oldRating, err := service.repository.GetByUserAndTarget(&rating.UserID, &rating.TargetID, rating.TargetType)
	if err != nil {
		return err
	}

	err = service.orchestrator.Start(rating, oldRating)
	if err != nil {
		return err
	}
	return nil
}

func (service *RatingService) Delete(rating *Rating) error {
	return service.repository.Delete(rating)
}

func (service *RatingService) GetAccommodationAverage(accommodationId *uuid.UUID) (float64, error) {
	return service.repository.GetTargetAverage(accommodationId, 0)
}

func (service *RatingService) GetHostAverage(hostId *uuid.UUID) (float64, error) {
	return service.repository.GetTargetAverage(hostId, 1)
}

func (service *RatingService) GetAccommodationRatings() (*[]Rating, error) {
	return service.repository.GetByTargetType(0)
}

func (service *RatingService) GetHostRatings() (*[]Rating, error) {
	return service.repository.GetByTargetType(1)
}

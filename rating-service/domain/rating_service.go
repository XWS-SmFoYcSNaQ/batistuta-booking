package domain

import (
	"errors"
	"github.com/google/uuid"
)

type RatingService struct {
	repository         RatingRepository
	createOrchestrator *CreateRatingOrchestrator
	deleteOrchestrator *DeleteRatingOrchestrator
}

func NewRatingService(repository *RatingRepository, createOrchestrator *CreateRatingOrchestrator, deleteOrchestrator *DeleteRatingOrchestrator) *RatingService {
	return &RatingService{
		repository:         *repository,
		createOrchestrator: createOrchestrator,
		deleteOrchestrator: deleteOrchestrator,
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

	err = service.createOrchestrator.Start(rating, oldRating)
	if err != nil {
		return err
	}
	return nil
}

func (service *RatingService) DeleteRating(id *uuid.UUID) error {
	var oldRating *Rating = nil
	oldRating, err := service.repository.GetById(id)
	if err != nil {
		return err
	}
	if oldRating == nil {
		return errors.New("rating not found")
	}
	err = service.deleteOrchestrator.Start(id, oldRating)
	if err != nil {
		return err
	}
	return nil
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

func (service *RatingService) Delete(rating *Rating) error {
	return service.repository.Delete(rating)
}

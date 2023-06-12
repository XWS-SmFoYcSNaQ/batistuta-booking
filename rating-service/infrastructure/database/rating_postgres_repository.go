package database

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingPostgresRepository struct {
	db *gorm.DB
}

func NewRatingPostgresRepository(db *gorm.DB) (domain.RatingRepository, error) {
	err := db.AutoMigrate(&domain.Rating{})
	if err != nil {
		return nil, err
	}
	return &RatingPostgresRepository{
		db: db,
	}, nil
}

func (store *RatingPostgresRepository) Insert(rating *domain.Rating) error {
	result := store.db.Create(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *RatingPostgresRepository) Update(rating *domain.Rating) error {
	result := store.db.Where(&domain.Rating{ID: rating.ID}).Updates(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *RatingPostgresRepository) GetAll() (*[]domain.Rating, error) {
	var ratings []domain.Rating
	result := store.db.Find(&ratings)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ratings, nil
}

func (store *RatingPostgresRepository) DeleteAll() {
	store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&domain.Rating{})
}

func (store *RatingPostgresRepository) Delete(rating *domain.Rating) error {
	result := store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *RatingPostgresRepository) GetByUserAndTarget(userId *uuid.UUID, targetId *uuid.UUID, targetType uint32) (*domain.Rating, error) {
	rating := domain.Rating{}
	result := store.db.Where(&domain.Rating{UserID: *userId, TargetID: *targetId, TargetType: targetType}).Find(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	if rating.ID == uuid.Nil {
		return nil, nil
	}
	return &rating, nil
}

package database

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
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

func (store *RatingPostgresRepository) Insert(product *domain.Rating) error {
	result := store.db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *RatingPostgresRepository) GetAll() (*[]domain.Rating, error) {
	var products []domain.Rating
	result := store.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
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

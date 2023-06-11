package services

import (
	"database/sql"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
)

type RatingService struct {
	DB *sql.DB
}

func (s RatingService) Create(rating *model.Rating) error {
	stmt, err := s.DB.Prepare("INSERT INTO Rating VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(rating.ID, rating.AccommodationId)
	if err != nil {
		return err
	}

	return nil
}

func (s RatingService) Delete(rating *model.Rating) error {
	stmt, err := s.DB.Prepare("DELETE FROM Rating WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(rating.ID)
	if err != nil {
		return err
	}

	return nil
}

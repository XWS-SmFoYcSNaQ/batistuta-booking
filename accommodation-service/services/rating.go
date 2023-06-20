package services

import (
	"database/sql"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
)

type RatingService struct {
	DB *sql.DB
}

func (s RatingService) Create(rating *model.Rating) error {
	stmt, err := s.DB.Prepare("INSERT INTO Rating VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(rating.ID, rating.AccommodationId, rating.UserID, rating.Value)
	if err != nil {
		return err
	}

	return nil
}

func (s RatingService) Update(rating *model.Rating) error {
	stmt, err := s.DB.Prepare("UPDATE Rating SET accommodation_id = $1, user_id = $2, value_ = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(rating.AccommodationId, rating.UserID, rating.Value, rating.ID)
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

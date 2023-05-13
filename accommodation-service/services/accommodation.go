package services

import (
	"accommodation_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type AccommodationService struct {
	DB *sql.DB
}

func (s AccommodationService) GetAll() ([]*model.Accommodation, error) {
	errorMessage := "error while fetching accommodations"
	rows, err := s.DB.Query("SELECT * FROM Accommodation")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	var accommodations []*model.Accommodation
	for rows.Next() {
		var p model.Accommodation
		err := rows.Scan(&p.ID, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		accommodations = append(accommodations, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New(errorMessage)
	}

	return accommodations, nil
}

func (s AccommodationService) Create(a *model.Accommodation) (uuid.UUID, error) {
	errorMessage := "error while creating accommodation"
	stmt, err := s.DB.Prepare("INSERT INTO Accommodation VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, a.Name, a.Benefits, a.MinGuests, a.MaxGuests, a.BasePrice)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}

	return id, nil
}

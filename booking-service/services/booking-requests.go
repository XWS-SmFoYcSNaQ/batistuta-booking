package services

import (
	"booking_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type BookingRequestsService struct {
	DB *sql.DB
}

func (s BookingRequestsService) GetAll() ([]*model.BookingRequest, error) {
	errorMessage := "error while fetching accommodations"
	rows, err := s.DB.Query("SELECT * FROM BookingRequest")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	var requests []*model.BookingRequest
	for rows.Next() {
		var p model.BookingRequest
		err := rows.Scan(&p.ID, &p.AccommodationId, &p.StartDate, &p.EndDate, &p.NumberOfGuests, &p.UserId)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		requests = append(requests, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New(errorMessage)
	}

	return requests, nil
}

func (s BookingRequestsService) MakeBookingRequest(r *model.BookingRequest) (uuid.UUID, error) {
	errorMessage := "error while creating accommodation"
	stmt, err := s.DB.Prepare("INSERT INTO BookingRequest VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, r.AccommodationId, r.StartDate, r.EndDate, r.NumberOfGuests, r.UserId)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}

	return id, nil
}

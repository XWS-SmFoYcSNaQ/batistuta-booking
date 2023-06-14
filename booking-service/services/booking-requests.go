package services

import (
	"booking_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type BookingRequestsService struct {
	DB *sql.DB
}

func (s BookingRequestsService) GetAll() ([]*model.BookingRequest, error) {
	errorMessage := "error while fetching booking requests"
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

func (s BookingRequestsService) GetAllByUserId(userId string) ([]*model.BookingRequest, error) {
	var rows *sql.Rows
	stmt, err := s.DB.Prepare("SELECT * FROM BookingRequest WHERE user_id = $1")
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.BookingRequest
	for rows.Next() {
		var p model.BookingRequest
		err := rows.Scan(&p.ID, &p.AccommodationId, &p.StartDate, &p.EndDate, &p.NumberOfGuests, &p.UserId)
		if err != nil {
			return nil, err
		}
		requests = append(requests, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
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

func (s BookingRequestsService) DeleteBookingRequest(id string) error {
	errorMessage := "error while deleting booking request"

	stmt, err := s.DB.Prepare("DELETE FROM BookingRequest WHERE id=$1")
	if err != nil {
		return errors.New(errorMessage)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return errors.New(errorMessage)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New(errorMessage)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // return error if no rows were affected
	}

	return nil
}

func (s BookingRequestsService) ConfirmReservation(id string) error {
	errorMessage := "error while confirming reservation"

	// Find the booking request
	row := s.DB.QueryRow("SELECT * FROM BookingRequest WHERE id = $1", id)
	var request model.BookingRequest
	err := row.Scan(&request.ID, &request.AccommodationId, &request.StartDate, &request.EndDate, &request.NumberOfGuests, &request.UserId)
	if err != nil {
		return errors.New(errorMessage)
	}

	// Create the reservation
	reservation := model.Reservation{
		ID:              uuid.New(),
		AccommodationId: request.AccommodationId,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		NumberOfGuests:  request.NumberOfGuests,
		UserId:          request.UserId,
	}

	// Insert the reservation into the database
	_, err = s.DB.Exec("INSERT INTO Reservation (id, accommodation_id, start_date, end_date, number_of_guests, user_id) VALUES ($1, $2, $3, $4, $5, $6)",
		reservation.ID, reservation.AccommodationId, reservation.StartDate, reservation.EndDate, reservation.NumberOfGuests, reservation.UserId)
	if err != nil {
		return errors.New(errorMessage)
	}

	// Delete the booking request
	_, err = s.DB.Exec("DELETE FROM BookingRequest WHERE id = $1", id)
	if err != nil {
		return errors.New(errorMessage)
	}

	return nil
}

func (s BookingRequestsService) GetAllReservationsForUser(userId string) ([]*model.Reservation, error) {
	var rows *sql.Rows
	stmt, err := s.DB.Prepare("SELECT * FROM Reservation WHERE user_id = $1")
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.Reservation
	for rows.Next() {
		var p model.Reservation
		err := rows.Scan(&p.ID, &p.AccommodationId, &p.StartDate, &p.EndDate, &p.NumberOfGuests, &p.UserId)
		if err != nil {
			return nil, err
		}
		requests = append(requests, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

func (s BookingRequestsService) DeleteReservation(reservationId string) error {
	errorMessage := "error while deleting reservation"

	// Find the reservation
	row := s.DB.QueryRow("SELECT * FROM Reservation WHERE id = $1", reservationId)
	var reservation model.Reservation
	err := row.Scan(&reservation.ID, &reservation.AccommodationId, &reservation.StartDate, &reservation.EndDate, &reservation.NumberOfGuests, &reservation.UserId)
	if err != nil {
		return errors.New(errorMessage)
	}

	// Check if the reservation can be cancelled
	startDate, err := time.Parse("2006-01-02 15:04:05-07", reservation.StartDate)
	if err != nil {
		return errors.New(errorMessage)
	}

	// Calculate the difference between the start date and today's date
	daysUntilStart := int(startDate.Sub(time.Now()).Hours() / 24)

	// Allow cancellation only if there is at least one day until the start date
	if daysUntilStart <= 1 {
		return errors.New("Reservation cannot be cancelled because it is too late to cancel it!")
	}

	// Delete the reservation
	_, err = s.DB.Exec("DELETE FROM Reservation WHERE id = $1", reservationId)
	if err != nil {
		return errors.New(errorMessage)
	}

	return nil

}

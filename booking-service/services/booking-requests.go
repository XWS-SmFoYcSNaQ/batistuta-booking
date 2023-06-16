package services

import (
	"booking_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
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
	_, err = stmt.Exec(id, r.AccommodationId, r.StartDate[:len(r.StartDate)-6], r.EndDate[:len(r.EndDate)-6], r.NumberOfGuests, r.UserId)
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

	// Delete overlapping booking requests
	_, err = s.DB.Exec("DELETE FROM BookingRequest WHERE accommodation_id = $1 AND TO_TIMESTAMP(start_date, 'YYYY-MM-DD HH24:MI:SS') <= TO_TIMESTAMP($2, 'YYYY-MM-DD HH24:MI:SS') AND TO_TIMESTAMP(end_date, 'YYYY-MM-DD HH24:MI:SS') >= TO_TIMESTAMP($3, 'YYYY-MM-DD HH24:MI:SS')",
		reservation.AccommodationId, reservation.EndDate, reservation.StartDate)
	if err != nil {
		return errors.New("Error while deleting other requests")
	}

	return nil
}

func (s BookingRequestsService) GetAllReservationsForUser(userId string) ([]*model.Reservation, error) {
	var rows *sql.Rows
	stmt, err := s.DB.Prepare("SELECT * FROM Reservation WHERE user_id = $1 AND is_active = true")
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
		err := rows.Scan(&p.ID, &p.AccommodationId, &p.StartDate, &p.EndDate, &p.NumberOfGuests, &p.UserId, &p.IsActive)
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
	err := row.Scan(&reservation.ID, &reservation.AccommodationId, &reservation.StartDate, &reservation.EndDate, &reservation.NumberOfGuests, &reservation.UserId, &reservation.IsActive)
	if err != nil {
		return errors.New("error while selecting all Reservations")
	}

	// Check if the reservation can be cancelled
	startDate, err := time.Parse("2006-01-02 15:04:05", reservation.StartDate)
	if err != nil {
		return errors.New("error while checking date")
	}

	// Calculate the difference between the start date and today's date
	daysUntilStart := int(startDate.Sub(time.Now()).Hours() / 24)

	// Allow cancellation only if there is at least one day until the start date
	if daysUntilStart <= 1 {
		return errors.New("Reservation cannot be cancelled because it is too late to cancel it!")
	}

	// Update the reservation to set is_active to false
	_, err = s.DB.Exec("UPDATE Reservation SET is_active = false WHERE id = $1", reservationId)
	if err != nil {
		return errors.New(errorMessage)
	}

	return nil

}

func (s BookingRequestsService) GetReservationsForAccommodationIDs(accommodationIDs []string) ([]*model.Reservation, error) {
	// Convert accommodationIDs to []interface{}
	placeholders := make([]string, len(accommodationIDs))
	values := make([]interface{}, len(accommodationIDs))
	for i, id := range accommodationIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		values[i] = id
	}

	// Construct the query string with placeholders for the accommodation IDs
	query := "SELECT * FROM Reservation WHERE is_active = true AND accommodation_id IN (" + strings.Join(placeholders, ",") + ")"

	// Prepare the query statement
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query with the accommodation IDs as arguments
	rows, err := stmt.Query(values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Fetch the reservations
	var reservations []*model.Reservation
	for rows.Next() {
		var r model.Reservation
		err := rows.Scan(&r.ID, &r.AccommodationId, &r.StartDate, &r.EndDate, &r.NumberOfGuests, &r.UserId, &r.IsActive)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, &r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (s BookingRequestsService) GetReservationRequestsForAccommodationIDs(accommodationIDs []string) ([]*model.BookingRequest, error) {
	// Convert accommodationIDs to []interface{}
	placeholders := make([]string, len(accommodationIDs))
	values := make([]interface{}, len(accommodationIDs))
	for i, id := range accommodationIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		values[i] = id
	}

	// Construct the query string with placeholders for the accommodation IDs
	query := "SELECT * FROM BookingRequest WHERE accommodation_id IN (" + strings.Join(placeholders, ",") + ")"

	// Prepare the query statement
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query with the accommodation IDs as arguments
	rows, err := stmt.Query(values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Fetch the reservations
	var reservations []*model.BookingRequest
	for rows.Next() {
		var r model.BookingRequest
		err := rows.Scan(&r.ID, &r.AccommodationId, &r.StartDate, &r.EndDate, &r.NumberOfGuests, &r.UserId)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, &r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (s BookingRequestsService) GetNumberOfCanceledReservationsForGuest(userId string) int32 {
	var count int32

	row := s.DB.QueryRow("SELECT COUNT(*) FROM Reservation WHERE user_id = $1 AND is_active = false", userId)
	err := row.Scan(&count)
	if err != nil {
		// Handle the error, such as returning a default value or logging the error
		return 0
	}

	return count
}

func (s BookingRequestsService) IsTheCancellationRateLessThanFive(accommodationIDs []string) (bool, error) {
	// Convert accommodationIDs to []interface{}
	placeholders := make([]string, len(accommodationIDs))
	values := make([]interface{}, len(accommodationIDs))
	for i, id := range accommodationIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		values[i] = id
	}

	totalReservationsQuery := `
	SELECT COUNT(*) 
	FROM Reservation AS r 
	WHERE r.accommodation_id IN (` + strings.Join(placeholders, ",") + `)
`

	// Query to count the number of canceled reservations for the host's accommodations
	canceledReservationsQuery := `
		SELECT COUNT(*) 
		FROM Reservation AS r 
		WHERE r.is_active = false AND r.accommodation_id IN (` + strings.Join(placeholders, ",") + `)
	`

	// Execute the queries
	var totalReservations int
	var canceledReservations int
	err := s.DB.QueryRow(totalReservationsQuery, values...).Scan(&totalReservations)
	if err != nil {
		return false, err
	}

	err = s.DB.QueryRow(canceledReservationsQuery, values...).Scan(&canceledReservations)
	if err != nil {
		return false, err
	}

	// Calculate the cancellation rate
	cancellationRate := float64(canceledReservations) / float64(totalReservations) * 100

	// Check if the cancellation rate is less than 5%
	if cancellationRate < 5 {
		return true, nil
	}

	return false, nil
}

func (s BookingRequestsService) HasAtLeastFivePastReservations(accommodationIDs []string) (bool, error) {
	// Convert accommodationIDs to []interface{}
	placeholders := make([]string, len(accommodationIDs))
	values := make([]interface{}, len(accommodationIDs))
	for i, id := range accommodationIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		values[i] = id
	}

	// Query to count the number of past reservations for the host
	pastReservationsQuery := `
		SELECT COUNT(*) 
		FROM Reservation AS r 
		WHERE TO_TIMESTAMP(end_date, 'YYYY-MM-DD HH24:MI:SS') < CURRENT_DATE AND r.accommodation_id IN (` + strings.Join(placeholders, ",") + `)
	`

	// Execute the query
	var count int
	err := s.DB.QueryRow(pastReservationsQuery, values...).Scan(&count)
	if err != nil {
		return false, err
	}

	// Check if the count is at least 5
	if count >= 5 {
		return true, nil
	}

	return false, nil

}

func (s BookingRequestsService) IsTotalReservationDurationGreaterThanFiftyDays(accommodationIDs []string) (bool, error) {
	// Convert accommodationIDs to []interface{}
	if len(accommodationIDs) == 0 {
		return false, nil
	}
	placeholders := make([]string, len(accommodationIDs))
	values := make([]interface{}, len(accommodationIDs))
	for i, id := range accommodationIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		values[i] = id
	}

	// Query to calculate the total duration of all reservations
	totalDurationQuery := `
		SELECT COALESCE(SUM(DATE_PART('day', TO_TIMESTAMP(end_date, 'YYYY-MM-DD HH24:MI:SS') - TO_TIMESTAMP(start_date, 'YYYY-MM-DD HH24:MI:SS'))), 0) AS total_duration
		FROM Reservation AS r
		WHERE r.accommodation_id IN (` + strings.Join(placeholders, ",") + `)
	`

	// Execute the query
	var totalDuration float64
	err := s.DB.QueryRow(totalDurationQuery, values...).Scan(&totalDuration)
	if err != nil {
		return false, err
	}

	// Check if the total duration is greater than 50 days
	if totalDuration > 50 {
		return true, nil
	}

	return false, nil
}

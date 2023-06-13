package model

import "github.com/google/uuid"

type Reservation struct {
	ID              uuid.UUID
	AccommodationId string
	StartDate       string
	EndDate         string
	NumberOfGuests  int
	UserId          string
}

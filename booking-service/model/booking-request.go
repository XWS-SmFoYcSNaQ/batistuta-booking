package model

import "github.com/google/uuid"

type BookingRequest struct {
	ID              uuid.UUID
	AccommodationId string
	StartDate       string
	EndDate         string
	NumberOfGuests  int
	UserId          string
}

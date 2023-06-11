package model

import "github.com/google/uuid"

type Rating struct {
	ID              uuid.UUID `json:"id"`
	AccommodationId uuid.UUID `json:"accommodation_id"`
}

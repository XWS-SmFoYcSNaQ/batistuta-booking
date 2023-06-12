package model

import "github.com/google/uuid"

type Rating struct {
	ID              uuid.UUID `json:"id"`
	AccommodationId uuid.UUID `json:"accommodation_id"`
	UserID          uuid.UUID `json:"user_id"`
	Value           uint32    `json:"value_"`
}

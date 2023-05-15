package model

import (
	"github.com/google/uuid"
	"time"
)

type Discount struct {
	ID              uuid.UUID `json:"id"`
	Start           time.Time `json:"d_start"`
	End             time.Time `json:"d_end"`
	AccommodationId uuid.UUID `json:"accommodation_id"`
	UserId          uuid.UUID `json:"user_id"`
	Discount        float64   `json:"discount"`
}

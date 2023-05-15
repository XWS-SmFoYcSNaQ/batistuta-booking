package model

import (
	"github.com/google/uuid"
	"time"
)

type Period struct {
	ID              uuid.UUID `json:"id"`
	Start           time.Time `json:"p_start"`
	End             time.Time `json:"p_end"`
	AccommodationId uuid.UUID `json:"accommodation_id"`
	UserId          uuid.UUID `json:"user_id"`
	Guests          int       `json:"guests"`
}

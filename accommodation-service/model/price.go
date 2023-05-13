package model

import (
	"github.com/google/uuid"
	"time"
)

type Price struct {
	ID              uuid.UUID
	Start           time.Time
	End             time.Time
	AccommodationId uuid.UUID
	UserId          uuid.UUID
	Price           float64
}

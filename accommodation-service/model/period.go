package model

import (
	"github.com/google/uuid"
	"time"
)

type Period struct {
	ID              uuid.UUID
	Start           time.Time
	End             time.Time
	AccommodationId uuid.UUID
	UserId          uuid.UUID
}

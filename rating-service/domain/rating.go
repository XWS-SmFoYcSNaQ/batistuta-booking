package domain

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	UserID       uuid.UUID
	TargetID     uuid.UUID
	Value        uint32
	LastModified time.Time

	//accommodation:0
	//host:			1
	TargetType uint32
}

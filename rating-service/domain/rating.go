package domain

import "github.com/google/uuid"

type Rating struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	UserID   uuid.UUID
	TargetID uuid.UUID

	//host:			0
	//accommodation:1
	TargetType uint32
	Value      uint32
}

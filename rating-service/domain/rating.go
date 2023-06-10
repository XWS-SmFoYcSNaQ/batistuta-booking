package domain

import "github.com/google/uuid"

type Rating struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	UserID   uuid.UUID
	TargetID uuid.UUID
	Value    uint32

	//accommodation:0
	//host:			1
	TargetType uint32
}

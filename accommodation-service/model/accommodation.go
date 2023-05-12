package model

import "github.com/google/uuid"

type Accommodation struct {
	ID        uuid.UUID
	Name      string
	Benefits  string
	MinGuests int
	MaxGuests int
}

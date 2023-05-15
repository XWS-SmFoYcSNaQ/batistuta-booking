package model

import "github.com/google/uuid"

type Accommodation struct {
	ID        uuid.UUID
	HostId    uuid.UUID
	Name      string
	Benefits  string
	MinGuests int
	MaxGuests int
	BasePrice float64
	Periods   []*Period
	Discounts []*Discount
}

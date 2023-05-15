package model

import "github.com/google/uuid"

type Accommodation struct {
	ID        uuid.UUID `json:"id"`
	HostId    uuid.UUID `json:"host_id"`
	Name      string    `json:"name"`
	Benefits  string    `json:"benefits"`
	MinGuests int       `json:"min_guests"`
	MaxGuests int       `json:"max_guests"`
	BasePrice float64   `json:"base_price"`
	Location  string
	Periods   []*Period   `json:"periods"`
	Discounts []*Discount `json:"discounts"`
}

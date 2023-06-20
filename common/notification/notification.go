package notification

import "github.com/google/uuid"

type Message struct {
	Title      string
	Content    string
	Type       Type
	NotifierId uuid.UUID
	ActorId    uuid.UUID
}

type Type int8

const (
	ReservationRequestCreated Type = iota
	ReservationCancelled
	HostRated
	AccommodationRated
	HostFeaturedStatusChanged
	ReservationRequestResponded
)

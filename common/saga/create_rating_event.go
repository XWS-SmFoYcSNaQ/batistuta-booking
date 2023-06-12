package saga

import "github.com/google/uuid"

type RatingDetails struct {
	ID         uuid.UUID
	TargetID   uuid.UUID
	TargetType uint32
	UserID     uuid.UUID
	Value      uint32
}

//command

type CreateRatingCommandType int8

const (
	UpdateUser CreateRatingCommandType = iota
	UpdateAccommodation
	UpdateHost
	RollbackRating
	ConcludeRatingCreation
	UnknownCommand
)

type CreateRatingCommand struct {
	Rating RatingDetails
	Type   CreateRatingCommandType
}

//reply

type CreateRatingReplyType int8

const (
	UserUpdated CreateRatingReplyType = iota
	UserUpdateFailed
	AccommodationUpdated
	AccommodationUpdateFailed
	HostUpdated
	HostUpdateFailed
	RatingRolledBack
	RatingCreationConcluded
	UnknownReply
)

type CreateRatingReply struct {
	Rating RatingDetails
	Type   CreateRatingReplyType
}

package saga

import "github.com/google/uuid"

type RatingDetails struct {
	ID         uuid.UUID
	TargetID   uuid.UUID
	TargetType uint32
	UserID     uuid.UUID
	Value      uint32
	OldValue   *RatingDetails
}

//command

type CreateRatingCommandType int8

const (
	StartRatingCreation CreateRatingCommandType = iota
	UpdateUser
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
	CreationStarted CreateRatingReplyType = iota
	CreationFailed
	UserUpdated
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

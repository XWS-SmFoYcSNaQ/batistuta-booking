package create_rating

import (
	"github.com/google/uuid"
	"time"
)

type RatingDetails struct {
	ID           uuid.UUID
	TargetID     uuid.UUID
	TargetType   uint32
	UserID       uuid.UUID
	Value        uint32
	LastModified time.Time
	OldValue     *RatingDetails
}

//command

type CreateRatingCommandType int8

const (
	StartRatingCreation CreateRatingCommandType = iota
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

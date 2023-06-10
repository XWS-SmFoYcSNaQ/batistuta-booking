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
	Authenticate CreateRatingCommandType = iota
	CheckIfAccommodationEligible
	CheckIfHostEligible
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
	Authenticated CreateRatingReplyType = iota
	AuthenticationFailed
	RatingRolledBack
	AccommodationEligible
	AccommodationNotEligible
	HostEligible
	HostNotEligible
	RatingCreationConcluded
	UnknownReply
)

type CreateRatingReply struct {
	Rating RatingDetails
	Type   CreateRatingReplyType
}

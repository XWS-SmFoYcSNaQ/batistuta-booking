package delete_rating

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/create_rating"
	"github.com/google/uuid"
)

type RatingDetails struct {
	ID       uuid.UUID
	OldValue *create_rating.RatingDetails
}

//command

type DeleteRatingCommandType int8

const (
	StartRatingDeletion DeleteRatingCommandType = iota
	UpdateUser
	UpdateAccommodation
	UpdateHost
	RollbackRating
	ConcludeRatingDeletion
	UnknownCommand
)

type DeleteRatingCommand struct {
	Rating RatingDetails
	Type   DeleteRatingCommandType
}

//reply

type DeleteRatingReplyType int8

const (
	DeletionStarted DeleteRatingReplyType = iota
	DeletionFailed
	UserUpdated
	UserUpdateFailed
	AccommodationUpdated
	AccommodationUpdateFailed
	HostUpdated
	HostUpdateFailed
	RatingRolledBack
	RatingDeletionConcluded
	UnknownReply
)

type DeleteRatingReply struct {
	Rating RatingDetails
	Type   DeleteRatingReplyType
}

package domain

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/create_rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/delete_rating"
	"github.com/google/uuid"
)

type DeleteRatingOrchestrator struct {
	commandPublisher *messaging.Publisher
	replySubscriber  *messaging.Subscriber
}

func NewDeleteRatingOrchestrator(publisher *messaging.Publisher, subscriber *messaging.Subscriber) (*DeleteRatingOrchestrator, error) {
	o := &DeleteRatingOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := (*o.replySubscriber).Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeleteRatingOrchestrator) Start(id *uuid.UUID, oldValue *Rating) error {
	event := &delete_rating.DeleteRatingCommand{
		Type: delete_rating.StartRatingDeletion,
		Rating: delete_rating.RatingDetails{
			ID: *id,
			OldValue: &create_rating.RatingDetails{
				ID:         oldValue.ID,
				TargetID:   oldValue.TargetID,
				TargetType: oldValue.TargetType,
				UserID:     oldValue.UserID,
				Value:      oldValue.Value,
			},
		},
	}
	return (*o.commandPublisher).Publish(event)
}

func (o *DeleteRatingOrchestrator) handle(reply *delete_rating.DeleteRatingReply) {
	command := delete_rating.DeleteRatingCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply)
	if command.Type != delete_rating.UnknownCommand {
		_ = (*o.commandPublisher).Publish(command)
	}
}

func (o *DeleteRatingOrchestrator) nextCommandType(reply *delete_rating.DeleteRatingReply) delete_rating.DeleteRatingCommandType {
	switch (*reply).Type {
	case delete_rating.DeletionStarted:
		return delete_rating.UpdateUser
	case delete_rating.DeletionFailed:
		return delete_rating.RollbackRating
	case delete_rating.UserUpdated:
		if (*reply).Rating.OldValue.TargetType == 1 {
			return delete_rating.UpdateHost
		}
		return delete_rating.UpdateAccommodation
	case delete_rating.UserUpdateFailed:
		return delete_rating.RollbackRating
	case delete_rating.AccommodationUpdated:
		return delete_rating.ConcludeRatingDeletion
	case delete_rating.AccommodationUpdateFailed:
		return delete_rating.RollbackRating
	case delete_rating.HostUpdated:
		return delete_rating.ConcludeRatingDeletion
	case delete_rating.HostUpdateFailed:
		return delete_rating.RollbackRating
	default:
		return delete_rating.UnknownCommand
	}
}

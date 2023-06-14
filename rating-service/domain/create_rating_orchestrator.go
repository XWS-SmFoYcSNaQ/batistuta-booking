package domain

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/create_rating"
)

type CreateRatingOrchestrator struct {
	commandPublisher *messaging.Publisher
	replySubscriber  *messaging.Subscriber
}

func NewCreateRatingOrchestrator(publisher *messaging.Publisher, subscriber *messaging.Subscriber) (*CreateRatingOrchestrator, error) {
	o := &CreateRatingOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := (*o.replySubscriber).Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateRatingOrchestrator) Start(rating *Rating, oldValue *Rating) error {
	event := &create_rating.CreateRatingCommand{
		Type: create_rating.StartRatingCreation,
		Rating: create_rating.RatingDetails{
			ID:         rating.ID,
			TargetID:   rating.TargetID,
			TargetType: rating.TargetType,
			UserID:     rating.UserID,
			Value:      rating.Value,
		},
	}
	if oldValue != nil {
		event.Rating.OldValue = &create_rating.RatingDetails{
			ID:         oldValue.ID,
			TargetID:   oldValue.TargetID,
			TargetType: oldValue.TargetType,
			UserID:     oldValue.UserID,
			Value:      oldValue.Value,
		}
	}
	return (*o.commandPublisher).Publish(event)
}

func (o *CreateRatingOrchestrator) handle(reply *create_rating.CreateRatingReply) {
	command := create_rating.CreateRatingCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply)
	if command.Type != create_rating.UnknownCommand {
		_ = (*o.commandPublisher).Publish(command)
	}
}

func (o *CreateRatingOrchestrator) nextCommandType(reply *create_rating.CreateRatingReply) create_rating.CreateRatingCommandType {
	switch (*reply).Type {
	case create_rating.CreationStarted:
		return create_rating.UpdateUser
	case create_rating.CreationFailed:
		return create_rating.RollbackRating
	case create_rating.UserUpdated:
		if (*reply).Rating.TargetType == 1 {
			return create_rating.UpdateHost
		}
		return create_rating.UpdateAccommodation
	case create_rating.UserUpdateFailed:
		return create_rating.RollbackRating
	case create_rating.AccommodationUpdated:
		return create_rating.ConcludeRatingCreation
	case create_rating.AccommodationUpdateFailed:
		return create_rating.RollbackRating
	case create_rating.HostUpdated:
		return create_rating.ConcludeRatingCreation
	case create_rating.HostUpdateFailed:
		return create_rating.RollbackRating
	default:
		return create_rating.UnknownCommand
	}
}

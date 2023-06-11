package domain

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga"
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

func (o *CreateRatingOrchestrator) Start(rating *Rating) error {
	event := &saga.CreateRatingCommand{
		Type: saga.UpdateUser,
		Rating: saga.RatingDetails{
			ID:         rating.ID,
			TargetID:   rating.TargetID,
			TargetType: rating.TargetType,
			UserID:     rating.UserID,
			Value:      rating.Value,
		},
	}
	return (*o.commandPublisher).Publish(event)
}

func (o *CreateRatingOrchestrator) handle(reply *saga.CreateRatingReply) {
	command := saga.CreateRatingCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply)
	if command.Type != saga.UnknownCommand {
		_ = (*o.commandPublisher).Publish(command)
	}
}

func (o *CreateRatingOrchestrator) nextCommandType(reply *saga.CreateRatingReply) saga.CreateRatingCommandType {
	switch (*reply).Type {
	case saga.UserUpdated:
		if (*reply).Rating.TargetType == 1 {
			return saga.UpdateHost
		}
		return saga.UpdateAccommodation
	case saga.UserUpdateFailed:
		return saga.RollbackRating
	case saga.AccommodationUpdated:
		return saga.ConcludeRatingCreation
	case saga.AccommodationUpdateFailed:
		return saga.RollbackRating
	case saga.HostUpdated:
		return saga.ConcludeRatingCreation
	case saga.HostUpdateFailed:
		return saga.RollbackRating
	default:
		return saga.UnknownCommand
	}
}

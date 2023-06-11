package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"log"
)

type CreateRatingCommandHandler struct {
	ratingService     *domain.RatingService
	replyPublisher    *messaging.Publisher
	commandSubscriber *messaging.Subscriber
}

func NewCreateRatingCommandHandler(ratingService *domain.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber) (*CreateRatingCommandHandler, error) {
	o := &CreateRatingCommandHandler{
		ratingService:     ratingService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := (*o.commandSubscriber).Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateRatingCommandHandler) handle(command *saga.CreateRatingCommand) {
	reply := saga.CreateRatingReply{Rating: command.Rating}
	switch command.Type {
	//temporary begin
	case saga.UpdateUser:
		reply.Type = saga.UserUpdated
	case saga.UpdateHost:
		reply.Type = saga.HostUpdated
	//temporary end

	case saga.RollbackRating:
		(*handler.ratingService).Delete(&domain.Rating{ID: command.Rating.ID})
		log.Println("RATING ROLLED BACK")
		reply.Type = saga.RatingRolledBack
	case saga.ConcludeRatingCreation:
		log.Println("RATING CREATED SUCCESSFULLY")
		reply.Type = saga.RatingCreationConcluded
	default:
		reply.Type = saga.UnknownReply
	}

	if reply.Type != saga.UnknownReply {
		(*handler.replyPublisher).Publish(reply)
	}
}

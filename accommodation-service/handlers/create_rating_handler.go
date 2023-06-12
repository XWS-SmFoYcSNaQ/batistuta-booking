package handlers

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga"
	"log"
)

type CreateRatingCommandHandler struct {
	ratingService     *services.RatingService
	replyPublisher    *messaging.Publisher
	commandSubscriber *messaging.Subscriber
}

func NewCreateRatingCommandHandler(ratingService *services.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber) (*CreateRatingCommandHandler, error) {
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
	case saga.UpdateAccommodation:
		//TODO: Check if user can rate the given accommodation
		r := model.Rating{ID: command.Rating.ID, AccommodationId: command.Rating.TargetID}
		err := handler.ratingService.Create(&r)
		if err != nil {
			log.Println(err)
			reply.Type = saga.AccommodationUpdateFailed
		} else {
			reply.Type = saga.AccommodationUpdated
		}
	case saga.RollbackRating:
		err := handler.ratingService.Delete(&model.Rating{ID: command.Rating.ID})
		if err != nil {
			log.Println(err)
		}
		reply.Type = saga.UnknownReply
	default:
		reply.Type = saga.UnknownReply
	}

	if reply.Type != saga.UnknownReply {
		(*handler.replyPublisher).Publish(reply)
	}
}

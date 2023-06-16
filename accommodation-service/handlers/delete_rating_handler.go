package handlers

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/delete_rating"
	"log"
)

type DeleteRatingCommandHandler struct {
	ratingService     *services.RatingService
	replyPublisher    *messaging.Publisher
	commandSubscriber *messaging.Subscriber
}

func NewDeleteRatingCommandHandler(ratingService *services.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber) (*DeleteRatingCommandHandler, error) {
	o := &DeleteRatingCommandHandler{
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

func (handler *DeleteRatingCommandHandler) handle(command *delete_rating.DeleteRatingCommand) {
	reply := delete_rating.DeleteRatingReply{Rating: command.Rating}
	switch command.Type {
	case delete_rating.UpdateAccommodation:
		err := handler.ratingService.Delete(&model.Rating{ID: command.Rating.ID})
		if err != nil {
			reply.Type = delete_rating.AccommodationUpdateFailed
		} else {
			reply.Type = delete_rating.AccommodationUpdated
		}
	case delete_rating.RollbackRating:
		oldValue := command.Rating.OldValue
		r := model.Rating{
			ID:              oldValue.ID,
			AccommodationId: oldValue.TargetID,
			UserID:          oldValue.UserID,
			Value:           oldValue.Value,
		}
		err := handler.ratingService.Create(&r)
		if err != nil {
			log.Println(err)
		}
		reply.Type = delete_rating.UnknownReply
	default:
		reply.Type = delete_rating.UnknownReply
	}

	if reply.Type != delete_rating.UnknownReply {
		(*handler.replyPublisher).Publish(reply)
	}
}

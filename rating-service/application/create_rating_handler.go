package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
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
	//reply := saga.CreateRatingReply{Rating: command.Rating}

}

package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/google/uuid"
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

	case saga.StartRatingCreation:
		oldValue := command.Rating.OldValue
		var err error
		r := domain.Rating{
			ID:         command.Rating.ID,
			UserID:     command.Rating.UserID,
			TargetID:   command.Rating.TargetID,
			Value:      command.Rating.Value,
			TargetType: command.Rating.TargetType,
		}
		if oldValue == nil {
			r.ID = uuid.New()
			reply.Rating.ID = r.ID
			err = (*handler.ratingService).Insert(&r)
		} else {
			r.ID = (*oldValue).ID
			reply.Rating.ID = r.ID
			err = (*handler.ratingService).Update(&r)
		}
		if err != nil {
			log.Println(err)
			reply.Type = saga.CreationFailed
		} else {
			reply.Type = saga.CreationStarted
		}
	case saga.RollbackRating:
		oldValue := command.Rating.OldValue
		if oldValue == nil {
			(*handler.ratingService).Delete(&domain.Rating{ID: command.Rating.ID})
		} else {
			(*handler.ratingService).Update(&domain.Rating{
				ID:         oldValue.ID,
				TargetID:   oldValue.TargetID,
				UserID:     oldValue.UserID,
				TargetType: oldValue.TargetType,
				Value:      oldValue.Value,
			})
		}
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

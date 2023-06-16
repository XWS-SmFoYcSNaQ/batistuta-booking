package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/create_rating"
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

func (handler *CreateRatingCommandHandler) handle(command *create_rating.CreateRatingCommand) {
	reply := create_rating.CreateRatingReply{Rating: command.Rating}
	switch command.Type {
	//temporary begin
	case create_rating.UpdateUser:
		reply.Type = create_rating.UserUpdated
	case create_rating.UpdateHost:
		reply.Type = create_rating.HostUpdated
	//temporary end

	case create_rating.StartRatingCreation:
		oldValue := command.Rating.OldValue
		var err error
		r := domain.Rating{
			ID:           command.Rating.ID,
			UserID:       command.Rating.UserID,
			TargetID:     command.Rating.TargetID,
			Value:        command.Rating.Value,
			TargetType:   command.Rating.TargetType,
			LastModified: command.Rating.LastModified,
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
			reply.Type = create_rating.CreationFailed
		} else {
			reply.Type = create_rating.CreationStarted
		}
	case create_rating.RollbackRating:
		oldValue := command.Rating.OldValue
		if oldValue == nil {
			(*handler.ratingService).Delete(&domain.Rating{ID: command.Rating.ID})
		} else {
			(*handler.ratingService).Update(&domain.Rating{
				ID:           oldValue.ID,
				TargetID:     oldValue.TargetID,
				UserID:       oldValue.UserID,
				TargetType:   oldValue.TargetType,
				Value:        oldValue.Value,
				LastModified: command.Rating.LastModified,
			})
		}
		log.Println("RATING ROLLED BACK")
		reply.Type = create_rating.RatingRolledBack
	case create_rating.ConcludeRatingCreation:
		log.Println("RATING CREATED SUCCESSFULLY")
		reply.Type = create_rating.RatingCreationConcluded
	default:
		reply.Type = create_rating.UnknownReply
	}

	if reply.Type != create_rating.UnknownReply {
		(*handler.replyPublisher).Publish(reply)
	}
}

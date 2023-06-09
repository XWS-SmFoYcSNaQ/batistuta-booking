package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/notification"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/saga/delete_rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"log"
)

type DeleteRatingCommandHandler struct {
	ratingService         *domain.RatingService
	replyPublisher        *messaging.Publisher
	commandSubscriber     *messaging.Subscriber
	notificationPublisher *messaging.Publisher
}

func NewDeleteRatingCommandHandler(ratingService *domain.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber, notificationPublisher *messaging.Publisher) (*DeleteRatingCommandHandler, error) {
	o := &DeleteRatingCommandHandler{
		ratingService:         ratingService,
		replyPublisher:        publisher,
		commandSubscriber:     subscriber,
		notificationPublisher: notificationPublisher,
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
	case delete_rating.StartRatingDeletion:
		err := (*handler.ratingService).Delete(&domain.Rating{ID: command.Rating.ID})
		if err != nil {
			reply.Type = delete_rating.DeletionFailed
		} else {
			reply.Type = delete_rating.DeletionStarted
		}
	case delete_rating.RollbackRating:
		oldValue := command.Rating.OldValue
		r := domain.Rating{
			ID:           oldValue.ID,
			UserID:       oldValue.UserID,
			TargetID:     oldValue.TargetID,
			Value:        oldValue.Value,
			TargetType:   oldValue.TargetType,
			LastModified: oldValue.LastModified,
		}
		err := (*handler.ratingService).Insert(&r)
		if err != nil {
			log.Println(err)
		}
		log.Println("RATING ROLLED BACK")
		reply.Type = delete_rating.RatingRolledBack
	case delete_rating.ConcludeRatingDeletion:
		log.Println("RATING DELETED SUCCESSFULLY")
		(*handler.notificationPublisher).Publish(&notification.Message{
			Title:      "Rating Created Successfully",
			Content:    "Rating Created Successfully",
			Type:       notification.AccommodationRated,
			NotifierId: (*(*command).Rating.OldValue).UserID,
		})
		reply.Type = delete_rating.RatingDeletionConcluded
	default:
		reply.Type = delete_rating.UnknownReply
	}

	if reply.Type != delete_rating.UnknownReply {
		(*handler.replyPublisher).Publish(reply)
	}
}

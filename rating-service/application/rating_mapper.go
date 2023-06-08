package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
)

func mapRating(r *domain.Rating) *rating.RatingDTO {
	return &rating.RatingDTO{
		Id:         r.ID.String(),
		UserId:     r.UserID.String(),
		TargetId:   r.TargetID.String(),
		TargetType: r.TargetType,
		Value:      r.Value,
	}
}

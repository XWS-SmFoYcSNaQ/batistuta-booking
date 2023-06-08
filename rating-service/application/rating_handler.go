package application

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
)

type RatingHandler struct {
	service *domain.RatingService
	rating.UnimplementedRatingServiceServer
}

func NewRatingHandler(service *domain.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (handler *RatingHandler) GetAll(ctx context.Context, request *rating.Empty) (*rating.RatingsList, error) {
	ratings, err := handler.service.GetAll()
	if err != nil || *ratings == nil {
		return nil, err
	}
	response := &rating.RatingsList{
		Data: []*rating.RatingDTO{},
	}
	for _, rating := range *ratings {
		current := mapRating(&rating)
		response.Data = append(response.Data, current)
	}
	return response, nil
}

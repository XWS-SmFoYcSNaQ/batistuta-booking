package application

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RatingHandler struct {
	service     *domain.RatingService
	authService *services.AuthService
	rating.UnimplementedRatingServiceServer
}

func NewRatingHandler(service *domain.RatingService, authService *services.AuthService) *RatingHandler {
	return &RatingHandler{
		service:     service,
		authService: authService,
	}
}

func (handler *RatingHandler) GetAllRatings(ctx context.Context, request *rating.Empty) (*rating.RatingsList, error) {
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

func (handler *RatingHandler) CreateAccommodationRating(ctx context.Context, request *rating.CreateAccommodationRatingDTO) (*rating.Empty, error) {
	res, err := handler.authService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	userId, err := uuid.Parse((*res).UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	targetId, err := uuid.Parse(request.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r := domain.Rating{
		TargetID: targetId,
		UserID:   userId,
		Value:    request.Value,
	}
	err = handler.service.CreateAccommodationRating(&r)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rating.Empty{}, nil
}

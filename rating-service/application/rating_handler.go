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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &rating.RatingsList{
		Data: []*rating.RatingDTO{},
	}
	for _, r := range *ratings {
		current := mapRating(&r)
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
		TargetID:   targetId,
		UserID:     userId,
		Value:      request.Value,
		TargetType: 0,
	}
	err = handler.service.CreateRating(&r)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rating.Empty{}, nil
}

func (handler *RatingHandler) CreateHostRating(ctx context.Context, request *rating.CreateHostRatingDTO) (*rating.Empty, error) {
	res, err := handler.authService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	userId, err := uuid.Parse((*res).UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	targetId, err := uuid.Parse(request.HostId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r := domain.Rating{
		TargetID:   targetId,
		UserID:     userId,
		Value:      request.Value,
		TargetType: 1,
	}
	err = handler.service.CreateRating(&r)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rating.Empty{}, nil
}

func (handler *RatingHandler) GetHostAverage(ctx context.Context, request *rating.IdMessage) (*rating.HostAverageDTO, error) {
	hostId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	average, err := handler.service.GetHostAverage(&hostId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rating.HostAverageDTO{
		HostId:  hostId.String(),
		Average: average,
	}, nil
}

func (handler *RatingHandler) GetHostRatings(ctx context.Context, request *rating.Empty) (*rating.RatingsList, error) {
	ratings, err := handler.service.GetHostRatings()
	if err != nil || *ratings == nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &rating.RatingsList{
		Data: []*rating.RatingDTO{},
	}
	for _, r := range *ratings {
		current := mapRating(&r)
		response.Data = append(response.Data, current)
	}
	return response, nil
}

func (handler *RatingHandler) Delete(ctx context.Context, request *rating.IdMessage) (*rating.Empty, error) {
	res, err := handler.authService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	_, err = uuid.Parse((*res).UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ratingId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = handler.service.DeleteRating(&ratingId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rating.Empty{}, nil
}

package controller

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"accommodation_service/services"
	"accommodation_service/utility"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationController struct {
	AccommodationService *services.AccommodationService
}

func (c AccommodationController) GetAll(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	accommodations, err := c.AccommodationService.GetAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.AccommodationSliceToDTOSlice(accommodations)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_GetAllAccommodations_Response{Data: r}, nil
}

func (c AccommodationController) Create(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	id, err := c.AccommodationService.Create(&model.Accommodation{
		Name:      request.Name,
		Benefits:  request.Benefits,
		MinGuests: int(request.MinGuests),
		MaxGuests: int(request.MaxGuests),
		BasePrice: request.BasePrice,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_CreateAccommodation_Response{Id: id.String()}, nil
}

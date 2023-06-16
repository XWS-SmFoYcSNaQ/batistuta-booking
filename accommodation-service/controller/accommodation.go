package controller

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/utility"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/auth"
	commonServices "github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationController struct {
	AccommodationService *services.AccommodationService
	PeriodService        *services.PeriodService
	DiscountService      *services.DiscountService
	AuthService          *commonServices.AuthService
}

func (c AccommodationController) GetAll(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	filters, err := utility.ExtractAccommodationFilters(request.Range, request.Benefits, request.Distinguished)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accommodations, err := c.AccommodationService.GetAll(filters)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.AccommodationSliceToDTOSlice(accommodations)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_GetAllAccommodations_Response{Data: r}, nil
}

func (c AccommodationController) GetMyAccommodations(ctx context.Context, request *accommodation.AM_GetMyAccommodations_Request) (*accommodation.AM_GetMyAccommodations_Response, error) {
	res, err := c.AuthService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	id, err := uuid.Parse((*res).UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accommodations, err := c.AccommodationService.GetAllByHostId(&id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	data, err := utility.AccommodationSliceToDTOSlice(accommodations)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetMyAccommodations_Response{Data: data}, nil
}

func (c AccommodationController) Create(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	res, err := c.AuthService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if res.UserRole != auth.UserRole_Host {
		return nil, status.Error(codes.Unauthenticated, "User is not a host")
	}
	hostId, err := uuid.Parse(res.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing host id")
	}

	id, err := c.AccommodationService.Create(&model.Accommodation{
		Name:      request.Name,
		HostId:    hostId,
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

func (c AccommodationController) GetById(ctx context.Context, request *accommodation.AM_GetAccommodation_Request) (*accommodation.AM_GetAccommodation_Response, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	a, err := c.AccommodationService.GetById(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return utility.AccommodationDetailsToDTO(a)
}

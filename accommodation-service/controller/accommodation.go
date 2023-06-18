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
	"strconv"
	"strings"
)

type AccommodationController struct {
	AccommodationService *services.AccommodationService
	PeriodService        *services.PeriodService
	DiscountService      *services.DiscountService
	AuthService          *commonServices.AuthService
}

func (c AccommodationController) GetAll(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	accommodations, err := c.AccommodationService.GetAll(uuid.Nil)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.AccommodationSliceToDTOSlice(accommodations)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_GetAllAccommodations_Response{Data: r}, nil
}

func (c AccommodationController) GetAllByHost(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	res, err := c.AuthService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	id, err := uuid.Parse((*res).UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accommodations, err := c.AccommodationService.GetAll(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	data, err := utility.AccommodationSliceToDTOSlice(accommodations)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAllAccommodations_Response{Data: data}, nil
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
		Name:                 request.Name,
		HostId:               hostId,
		Benefits:             request.Benefits,
		MinGuests:            int(request.MinGuests),
		MaxGuests:            int(request.MaxGuests),
		BasePrice:            request.BasePrice,
		Location:             request.Location,
		AutomaticReservation: request.AutomaticReservation,
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

func (c AccommodationController) SearchAccommodations(ctx context.Context, request *accommodation.AM_SearchAccommodations_Request) (*accommodation.AM_SearchAccommodations_Response, error) {
	if request.NumberOfGuests <= 0 {
		return nil, status.Error(codes.InvalidArgument, "number of guests must be greater than 0 ("+strconv.Itoa(int(request.NumberOfGuests))+")")
	}

	accommodations, err := c.AccommodationService.GetAccommodationSearchResults(&accommodation.AM_SearchAccommodations_Request{
		Start:          request.Start,
		End:            request.End,
		NumberOfGuests: request.NumberOfGuests,
		Location:       request.Location,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var availableAccommodations []*model.Accommodation
	for _, a := range accommodations {
		start, err := utility.ParseISOString(request.Start)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		end, err := utility.ParseISOString(request.End)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		available, err := c.PeriodService.IsAvailableForGivenInterval(a.ID, start, end)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if available && strings.Contains(a.Location, request.Location) {
			availableAccommodations = append(availableAccommodations, a)
		}
	}

	var res []*accommodation.AccommodationSearchResultDTO
	for _, d := range availableAccommodations {
		a := accommodation.AccommodationSearchResultDTO{
			Id:         d.ID.String(),
			Name:       d.Name,
			Benefits:   d.Benefits,
			MinGuests:  int32(d.MinGuests),
			MaxGuests:  int32(d.MaxGuests),
			BasePrice:  d.BasePrice,
			Location:   d.Location,
			TotalPrice: 0.0,
		}
		res = append(res, &a)
	}

	return &accommodation.AM_SearchAccommodations_Response{Data: res}, nil
}

func (c AccommodationController) GetAutomaticReservationValue(ctx context.Context, request *accommodation.AM_GetAutomaticReservation_Request) (*accommodation.AM_GetAutomaticReservation_Response, error) {
	resp, err := c.AccommodationService.GetAutomaticReservationValue(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAutomaticReservation_Response{AutomaticReservation: int32(resp)}, nil
}

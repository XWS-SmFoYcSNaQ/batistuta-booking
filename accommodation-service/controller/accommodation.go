package controller

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"accommodation_service/services"
	"accommodation_service/utility"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationController struct {
	AccommodationService *services.AccommodationService
	PeriodService        *services.PeriodService
	DiscountService      *services.DiscountService
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

func (c AccommodationController) GetByIdWithPeriods(ctx context.Context, request *accommodation.AM_GetAccommodationWithPeriods_Request) (*accommodation.AM_GetAccommodationWithPeriods_Response, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	a, err := c.AccommodationService.GetById(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while fetching accommodation")
	}

	periods, err := c.PeriodService.GetAllByAccommodation(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while fetching periods")
	}

	return utility.AccommodationWithPeriodsToDTO(a, periods)
}

func (c AccommodationController) GetByIdWithDiscounts(ctx context.Context, request *accommodation.AM_GetAccommodationWithDiscounts_Request) (*accommodation.AM_GetAccommodationWithDiscounts_Response, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	a, err := c.AccommodationService.GetById(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while fetching accommodation")
	}

	discounts, err := c.DiscountService.GetAllByAccommodation(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while fetching discounts")
	}

	return utility.AccommodationWithDiscountsToDTO(a, discounts)
}

func (c AccommodationController) SearchAccommodations(ctx context.Context, request *accommodation.AM_SearchAccommodations_Request) (*accommodation.AM_SearchAccommodations_Response, error) {
	if request.NumberOfGuests <= 0 {
		return nil, status.Error(codes.InvalidArgument, "number of guests must be greater than 0")
	}

	accommodations, err := c.AccommodationService.GetSearchedAccommodations(&accommodation.AM_SearchAccommodations_Request{
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
		if available {
			availableAccommodations = append(availableAccommodations, a)
		}
	}

	var res []*accommodation.SearchedAccommodationsDTO
	for _, d := range availableAccommodations {
		a := accommodation.SearchedAccommodationsDTO{
			Id:         d.ID.String(),
			Name:       d.Name,
			Benefits:   d.Benefits,
			MinGuests:  int32(d.MinGuests),
			MaxGuests:  int32(d.MaxGuests),
			BasePrice:  d.BasePrice,
			Location:   "nema jos",
			TotalPrice: 0.0,
		}
		res = append(res, &a)
	}

	return &accommodation.AM_SearchAccommodations_Response{Data: res}, nil
}

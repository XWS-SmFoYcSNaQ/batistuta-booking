package handlers

import (
	"accommodation_service/controller"
	"accommodation_service/proto/accommodation"
	"context"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	AccommodationController *controller.AccommodationController
	PeriodController        *controller.PeriodController
	DiscountController      *controller.DiscountController
}

func (h AccommodationHandler) GetAllAccommodations(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	return h.AccommodationController.GetAll(ctx, request)
}

func (h AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	return h.AccommodationController.Create(ctx, request)
}

func (h AccommodationHandler) GetAccommodationWithPeriods(ctx context.Context, request *accommodation.AM_GetAccommodationWithPeriods_Request) (*accommodation.AM_GetAccommodationWithPeriods_Response, error) {
	return h.AccommodationController.GetByIdWithPeriods(ctx, request)
}

func (h AccommodationHandler) GetAccommodationWithDiscounts(ctx context.Context, request *accommodation.AM_GetAccommodationWithDiscounts_Request) (*accommodation.AM_GetAccommodationWithDiscounts_Response, error) {
	return h.AccommodationController.GetByIdWithDiscounts(ctx, request)
}

func (h AccommodationHandler) GetAllPeriodsByAccommodation(ctx context.Context, request *accommodation.AM_GetAllPeriodsByAccommodation_Request) (*accommodation.AM_GetAllPeriodsByAccommodation_Response, error) {
	return h.PeriodController.GetAllByAccommodation(ctx, request)
}

func (h AccommodationHandler) CreatePeriod(ctx context.Context, request *accommodation.AM_CreatePeriod_Request) (*accommodation.AM_CreatePeriod_Response, error) {
	return h.PeriodController.Create(ctx, request)
}

func (h AccommodationHandler) GetAllDiscountsByAccommodation(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodation_Request) (*accommodation.AM_GetAllDiscountsByAccommodation_Response, error) {
	return h.DiscountController.GetAllByAccommodation(ctx, request)
}

func (h AccommodationHandler) GetAllDiscountsByAccommodationAndInterval(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Request) (*accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Response, error) {
	return h.DiscountController.GetAllByAccommodationAndInterval(ctx, request)
}

func (h AccommodationHandler) CreateDiscount(ctx context.Context, request *accommodation.AM_CreateDiscount_Request) (*accommodation.AM_CreateDiscount_Response, error) {
	return h.DiscountController.Create(ctx, request)
}

func (h AccommodationHandler) SearchAccommodations(ctx context.Context, request *accommodation.AM_SearchAccommodations_Request) (*accommodation.AM_SearchAccommodations_Response, error) {
	if request.NumberOfGuests <= 0 {
		return nil, status.Error(codes.InvalidArgument, "number of guests must be greater than 0")
	}

	accommodations, err := h.AccommodationService.GetSearchedAccommodations(&accommodation.AM_SearchAccommodations_Request{
		Start:          request.Start,
		End:            request.End,
		NumberOfGuests: request.NumberOfGuests,
		Location:       request.Location,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res []*accommodation.SearchedAccommodationsDTO
	for _, d := range accommodations {
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

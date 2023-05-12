package handlers

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"accommodation_service/services"
	"context"
	"database/sql"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	DB                   *sql.DB
	AccommodationService *services.AccommodationService
}

func (h AccommodationHandler) GetAllAccommodations(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	accommodations, err := h.AccommodationService.GetAll()
	if err != nil {
		return &accommodation.AM_GetAllAccommodations_Response{
			Error: err.Error(),
		}, nil
	}
	var res []*accommodation.AccommodationDTO
	for _, d := range accommodations {
		a := accommodation.AccommodationDTO{
			Id:        d.ID.String(),
			Name:      d.Name,
			Benefits:  d.Benefits,
			MinGuests: int32(d.MinGuests),
			MaxGuests: int32(d.MaxGuests),
		}
		res = append(res, &a)
	}

	return &accommodation.AM_GetAllAccommodations_Response{
		Data: res,
	}, nil
}

func (h AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	id, err := h.AccommodationService.Create(&model.Accommodation{Name: request.Name, Benefits: request.Benefits})
	if err != nil {
		return &accommodation.AM_CreateAccommodation_Response{
			Error: err.Error(),
		}, nil
	}
	return &accommodation.AM_CreateAccommodation_Response{
		Id: id.String(),
	}, nil
}

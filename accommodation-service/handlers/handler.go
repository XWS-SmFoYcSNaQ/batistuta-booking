package handlers

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"accommodation_service/services"
	"accommodation_service/utility"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	DB                   *sql.DB
	AccommodationService *services.AccommodationService
	PeriodService        *services.PeriodService
}

func (h AccommodationHandler) GetAllAccommodations(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	accommodations, err := h.AccommodationService.GetAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res []*accommodation.AccommodationDTO
	for _, d := range accommodations {
		a := accommodation.AccommodationDTO{
			Id:        d.ID.String(),
			Name:      d.Name,
			Benefits:  d.Benefits,
			MinGuests: int32(d.MinGuests),
			MaxGuests: int32(d.MaxGuests),
			BasePrice: d.BasePrice,
		}
		res = append(res, &a)
	}

	return &accommodation.AM_GetAllAccommodations_Response{Data: res}, nil
}

func (h AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	id, err := h.AccommodationService.Create(&model.Accommodation{
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

func (h AccommodationHandler) GetAllPeriodsByAccommodation(ctx context.Context, request *accommodation.AM_GetAllPeriodsByAccommodation_Request) (*accommodation.AM_GetAllPeriodsByAccommodation_Response, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	periods, err := h.PeriodService.GetAllByAccommodation(id)
	var res []*accommodation.PeriodDTO
	for _, d := range periods {
		userId := ""
		if d.UserId != uuid.Nil {
			userId = d.UserId.String()
		}
		p := accommodation.PeriodDTO{
			Id:              d.ID.String(),
			Start:           d.Start.String(),
			End:             d.End.String(),
			AccommodationId: d.AccommodationId.String(),
			UserId:          userId,
		}
		res = append(res, &p)
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAllPeriodsByAccommodation_Response{Data: res}, nil
}

func (h AccommodationHandler) CreatePeriod(ctx context.Context, request *accommodation.AM_CreatePeriod_Request) (*accommodation.AM_CreatePeriod_Response, error) {
	start, err := utility.ParseISOString(request.Start)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	end, err := utility.ParseISOString(request.End)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accommodationId, err := uuid.Parse(request.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	userId := uuid.Nil
	if request.UserId != "" {
		userId, err = uuid.Parse(request.UserId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	id, err := h.PeriodService.Create(&model.Period{
		Start:           start,
		End:             end,
		AccommodationId: accommodationId,
		UserId:          userId,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_CreatePeriod_Response{Id: id.String()}, nil
}

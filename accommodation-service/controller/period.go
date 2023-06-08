package controller

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/utility"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PeriodController struct {
	PeriodService *services.PeriodService
}

func (c PeriodController) GetAllByAccommodation(ctx context.Context, request *accommodation.AM_GetAllPeriodsByAccommodation_Request) (*accommodation.AM_GetAllPeriodsByAccommodation_Response, error) {
	id, err := uuid.Parse(request.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	periods, err := c.PeriodService.GetAllByAccommodation(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.PeriodSliceToDTOSlice(periods)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAllPeriodsByAccommodation_Response{Data: r}, nil
}

func (c PeriodController) Create(ctx context.Context, request *accommodation.AM_CreatePeriod_Request) (*accommodation.AM_CreatePeriod_Response, error) {
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
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	userId := uuid.Nil
	if request.UserId != "" {
		userId, err = uuid.Parse(request.UserId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Error while parsing user id")
		}
	}

	id, err := c.PeriodService.Create(&model.Period{
		Start:           start,
		End:             end,
		AccommodationId: accommodationId,
		UserId:          userId,
		Guests:          int(request.Guests),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_CreatePeriod_Response{Id: id.String()}, nil
}

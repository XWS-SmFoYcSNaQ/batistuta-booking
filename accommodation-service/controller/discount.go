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

type DiscountController struct {
	DiscountService *services.DiscountService
	AuthService     *commonServices.AuthService
}

func (c DiscountController) GetAllByAccommodation(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodation_Request) (*accommodation.AM_GetAllDiscountsByAccommodation_Response, error) {
	id, err := uuid.Parse(request.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Error while parsing accommodation id")
	}
	discounts, err := c.DiscountService.GetAllByAccommodation(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.DiscountSliceToDTOSlice(discounts)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAllDiscountsByAccommodation_Response{Data: r}, nil
}

func (c DiscountController) GetAllByAccommodationAndInterval(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Request) (*accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Response, error) {
	start, err := utility.ParseISOString(request.Start)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	end, err := utility.ParseISOString(request.End)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := uuid.Parse(request.AccommodationId)
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
	discounts, err := c.DiscountService.GetAllByAccommodationAndInterval(id, start, end, userId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	r, err := utility.DiscountSliceToDTOSlice(discounts)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Response{Data: r}, nil
}

func (c DiscountController) Create(ctx context.Context, request *accommodation.AM_CreateDiscount_Request) (*accommodation.AM_CreateDiscount_Response, error) {
	res, err := c.AuthService.ValidateToken(&ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if res.UserRole != auth.UserRole_Host {
		return nil, status.Error(codes.Unauthenticated, "User is not a host")
	}

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

	id, err := c.DiscountService.Create(&model.Discount{
		Start:           start,
		End:             end,
		AccommodationId: accommodationId,
		UserId:          userId,
		Discount:        request.Discount,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &accommodation.AM_CreateDiscount_Response{Id: id.String()}, nil
}

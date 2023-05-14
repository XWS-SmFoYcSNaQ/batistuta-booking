package handlers

import (
	"booking_service/model"
	"booking_service/proto/booking"
	"booking_service/services"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingHandler struct {
	booking.UnimplementedBookingServiceServer
	DB                    *sql.DB
	BookingRequestService *services.BookingRequestsService
}

func (h BookingHandler) GetAll(ctx context.Context, request *booking.AM_GetAllBookingRequests_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	bookingRequests, err := h.BookingRequestService.GetAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res []*booking.BookingRequestsDTO
	for _, d := range bookingRequests {
		a := booking.BookingRequestsDTO{
			Id:              d.ID.String(),
			AccommodationId: d.AccommodationId,
			StartDate:       d.StartDate,
			EndDate:         d.EndDate,
			NumberOfGuests:  int32(d.NumberOfGuests),
			UserId:          d.UserId,
		}
		res = append(res, &a)
	}

	return &booking.AM_GetAllBookingRequests_Response{Data: res}, nil
}

func (h BookingHandler) MakeBookingRequest(ctx context.Context, request *booking.AM_BookingRequest_Request) (*booking.AM_CreateBookingRequest_Response, error) {
	id, err := h.BookingRequestService.MakeBookingRequest(&model.BookingRequest{
		AccommodationId: request.AccommodationId,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		NumberOfGuests:  int(request.NumberOfGuests),
		UserId:          request.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.AM_CreateBookingRequest_Response{Id: id.String()}, nil
}

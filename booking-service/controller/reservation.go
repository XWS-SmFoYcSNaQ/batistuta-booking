package controller

import (
	"booking_service/model"
	"booking_service/proto/booking"
	"booking_service/services"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReservationController struct {
	BookingService *services.BookingRequestsService
}

func (c ReservationController) GetAll(ctx context.Context, request *booking.AM_GetAllBookingRequests_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	bookingRequests, err := c.BookingService.GetAll()
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

func (c ReservationController) GetAllByUserId(ctx context.Context, request *booking.AM_GetAllBookingRequestsByUserId_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	bookingRequests, err := c.BookingService.GetAllByUserId(request.Id)
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

func (c ReservationController) MakeBookingRequest(ctx context.Context, request *booking.AM_BookingRequest_Request) (*booking.AM_CreateBookingRequest_Response, error) {
	id, err := c.BookingService.MakeBookingRequest(&model.BookingRequest{
		AccommodationId: request.AccommodationId,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		UserId:          request.UserId,
		NumberOfGuests:  int(request.NumberOfGuests),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.AM_CreateBookingRequest_Response{Id: id.String()}, nil
}

func (c ReservationController) DeleteBookingRequest(ctx context.Context, request *booking.AM_DeleteBookingRequest_Request) (*booking.AM_DeleteBookingRequest_Response, error) {
	err := c.BookingService.DeleteBookingRequest(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.AM_DeleteBookingRequest_Response{}, nil
}

func (c ReservationController) ConfirmReservationRequest(ctx context.Context, request *booking.ReservationConfirm_Request) (*booking.EmptyMessage, error) {
	err := c.BookingService.ConfirmReservation(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.EmptyMessage{}, nil
}

func (c ReservationController) GetAllReservationsForGuest(ctx context.Context, request *booking.AllReservationsForGuest_Request) (*booking.AllReservationsForGuest_Response, error) {
	bookingRequests, err := c.BookingService.GetAllReservationsForUser(request.Id)
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

	return &booking.AllReservationsForGuest_Response{Data: res}, nil
}

func (c ReservationController) DeleteReservation(ctx context.Context, request *booking.DeleteReservation_Request) (*booking.EmptyMessage, error) {
	err := c.BookingService.DeleteReservation(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.EmptyMessage{}, nil
}

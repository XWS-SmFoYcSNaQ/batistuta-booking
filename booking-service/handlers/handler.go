package handlers

import (
	controller "booking_service/controller"
	"booking_service/proto/booking"
	"context"
)

type BookingHandler struct {
	booking.UnimplementedBookingServiceServer
	ReservationController *controller.ReservationController
}

func (h BookingHandler) GetAll(ctx context.Context, request *booking.AM_GetAllBookingRequests_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	return h.ReservationController.GetAll(ctx, request)
}

func (h BookingHandler) GetAllByUserId(ctx context.Context, request *booking.AM_GetAllBookingRequestsByUserId_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	return h.ReservationController.GetAllByUserId(ctx, request)
}

func (h BookingHandler) MakeBookingRequest(ctx context.Context, request *booking.AM_BookingRequest_Request) (*booking.AM_CreateBookingRequest_Response, error) {
	return h.ReservationController.MakeBookingRequest(ctx, request)
}

func (h BookingHandler) DeleteBookingRequest(ctx context.Context, request *booking.AM_DeleteBookingRequest_Request) (*booking.AM_DeleteBookingRequest_Response, error) {
	return h.ReservationController.DeleteBookingRequest(ctx, request)
}

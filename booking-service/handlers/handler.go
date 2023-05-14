package handlers

import (
	"booking_service/proto/booking"
	"context"
	"fmt"
)

type BookingHandler struct {
	booking.UnimplementedBookingServiceServer
}

func (h BookingHandler) GetAll(ctx context.Context, request *booking.BK_Request) (*booking.BK_Response, error) {
	return &booking.BK_Response{
		Message: fmt.Sprintf("Message %s!", request.Name),
	}, nil
}

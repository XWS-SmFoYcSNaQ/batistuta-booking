package handlers

import (
	"accommodation_service/proto/accommodation"
	"context"
	"fmt"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
}

func (h AccommodationHandler) GetAll(ctx context.Context, request *accommodation.AM_Request) (*accommodation.AM_Response, error) {
	return &accommodation.AM_Response{
		Greeting: fmt.Sprintln("Hello!"),
	}, nil
}

func (h AccommodationHandler) Create(ctx context.Context, request *accommodation.AM_Request) (*accommodation.AM_Response, error) {
	return &accommodation.AM_Response{
		Greeting: fmt.Sprintf("Create method! %s", request.Name),
	}, nil
}

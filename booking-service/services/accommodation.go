package services

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
)

type AccommodationService struct {
	AccomodationClient *accommodation.AccommodationServiceClient
}

func NewAccommodationService(accommodationClient *accommodation.AccommodationServiceClient) *AccommodationService {
	return &AccommodationService{
		AccomodationClient: accommodationClient,
	}
}

package mock

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/google/uuid"
)

var Ratings = []*domain.Rating{
	{
		ID:         uuid.New(),
		UserID:     uuid.New(),
		TargetID:   uuid.New(),
		TargetType: 0,
		Value:      5,
	},
}

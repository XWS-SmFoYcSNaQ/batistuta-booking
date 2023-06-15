package application

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/user"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
)

func MapRating(r *domain.Rating) *rating.RatingDTO {
	return &rating.RatingDTO{
		Id:           r.ID.String(),
		UserId:       r.UserID.String(),
		TargetId:     r.TargetID.String(),
		TargetType:   r.TargetType,
		Value:        r.Value,
		LastModified: r.LastModified.String(),
	}
}

func MapRatingDetailsList(ratings *[]domain.Rating, users []*user.User) []*rating.RatingDetailsDTO {
	var data []*rating.RatingDetailsDTO
	for _, r := range *ratings {
		u := *findUserById(r.UserID.String(), users)
		dto := &rating.RatingDetailsDTO{
			Id:            r.ID.String(),
			TargetId:      r.TargetID.String(),
			TargetType:    r.TargetType,
			Value:         r.Value,
			LastModified:  r.LastModified.String(),
			UserId:        r.UserID.String(),
			UserFirstName: u.FirstName,
			UserLastName:  u.LastName,
			UserEmail:     u.Email,
		}
		data = append(data, dto)
	}
	return data
}

func findUserById(id string, users []*user.User) *user.User {
	for _, u := range users {
		if (*u).Id == id {
			return u
		}
	}
	return &user.User{}
}

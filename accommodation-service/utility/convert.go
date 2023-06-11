package utility

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/google/uuid"
)

func AccommodationToDTO(d *model.Accommodation) (*accommodation.AccommodationDTO, error) {
	if d == nil {
		return nil, nil
	}
	return &accommodation.AccommodationDTO{
		Id:        d.ID.String(),
		Name:      d.Name,
		Benefits:  d.Benefits,
		MinGuests: int32(d.MinGuests),
		MaxGuests: int32(d.MaxGuests),
		BasePrice: d.BasePrice,
	}, nil
}

func AccommodationSliceToDTOSlice(data []*model.Accommodation) ([]*accommodation.AccommodationDTO, error) {
	if data == nil {
		return nil, nil
	}
	var res []*accommodation.AccommodationDTO
	for _, d := range data {
		a, err := AccommodationToDTO(d)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func PeriodToDTO(d *model.Period) (*accommodation.PeriodDTO, error) {
	if d == nil {
		return nil, nil
	}
	userId := ""
	if d.UserId != uuid.Nil {
		userId = d.UserId.String()
	}
	return &accommodation.PeriodDTO{
		Id:              d.ID.String(),
		Start:           d.Start.String(),
		End:             d.End.String(),
		AccommodationId: d.AccommodationId.String(),
		UserId:          userId,
	}, nil
}

func PeriodSliceToDTOSlice(data []*model.Period) ([]*accommodation.PeriodDTO, error) {
	if data == nil {
		return nil, nil
	}
	var res []*accommodation.PeriodDTO
	for _, d := range data {
		a, err := PeriodToDTO(d)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func DiscountToDTO(d *model.Discount) (*accommodation.DiscountDTO, error) {
	if d == nil {
		return nil, nil
	}
	userId := ""
	if d.UserId != uuid.Nil {
		userId = d.UserId.String()
	}
	return &accommodation.DiscountDTO{
		Id:              d.ID.String(),
		Start:           d.Start.String(),
		End:             d.End.String(),
		AccommodationId: d.AccommodationId.String(),
		UserId:          userId,
		Discount:        d.Discount,
	}, nil
}

func DiscountSliceToDTOSlice(data []*model.Discount) ([]*accommodation.DiscountDTO, error) {
	if data == nil {
		return nil, nil
	}
	var res []*accommodation.DiscountDTO
	for _, d := range data {
		a, err := DiscountToDTO(d)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func AccommodationDetailsToDTO(a *model.Accommodation) (*accommodation.AM_GetAccommodation_Response, error) {
	if a == nil {
		return nil, nil
	}
	periods, err := PeriodSliceToDTOSlice(a.Periods)
	discounts, err := DiscountSliceToDTOSlice(a.Discounts)
	if err != nil {
		return nil, err
	}
	return &accommodation.AM_GetAccommodation_Response{
		Id:        a.ID.String(),
		Name:      a.Name,
		Benefits:  a.Benefits,
		MinGuests: int32(a.MinGuests),
		MaxGuests: int32(a.MaxGuests),
		BasePrice: a.BasePrice,
		Periods:   periods,
		Discounts: discounts,
	}, nil
}

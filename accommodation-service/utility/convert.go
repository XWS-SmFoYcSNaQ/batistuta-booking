package utility

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"github.com/google/uuid"
)

func AccommodationToDTO(d *model.Accommodation) (*accommodation.AccommodationDTO, error) {
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

func AccommodationWithPeriodsToDTO(a *model.Accommodation, p []*model.Period) (*accommodation.AM_GetAccommodationWithPeriods_Response, error) {
	periods, err := PeriodSliceToDTOSlice(p)
	if err != nil {
		return nil, err
	}
	return &accommodation.AM_GetAccommodationWithPeriods_Response{
		Id:        a.ID.String(),
		Name:      a.Name,
		Benefits:  a.Benefits,
		MinGuests: int32(a.MinGuests),
		MaxGuests: int32(a.MaxGuests),
		BasePrice: a.BasePrice,
		Periods:   periods,
	}, nil
}

func AccommodationWithDiscountsToDTO(a *model.Accommodation, d []*model.Discount) (*accommodation.AM_GetAccommodationWithDiscounts_Response, error) {
	discounts, err := DiscountSliceToDTOSlice(d)
	if err != nil {
		return nil, err
	}
	return &accommodation.AM_GetAccommodationWithDiscounts_Response{
		Id:        a.ID.String(),
		Name:      a.Name,
		Benefits:  a.Benefits,
		MinGuests: int32(a.MinGuests),
		MaxGuests: int32(a.MaxGuests),
		BasePrice: a.BasePrice,
		Discounts: discounts,
	}, nil
}

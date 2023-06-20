package utility

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ParseISOString(isoString string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	parsedTime, err := time.Parse(layout, isoString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ExtractAccommodationFilters(priceRange string, benefits string, distinguished string) (*Filter, error) {
	var filter Filter
	r, err := extractRangeFilter(priceRange)
	if err != nil {
		return nil, err
	}
	filter.Range = r

	b, err := extractBenefitsFilter(benefits)
	if err != nil {
		return nil, err
	}
	filter.Benefits = *b

	filter.Distinguished = false
	if distinguished == "true" {
		filter.Distinguished = true
	}

	return &filter, nil
}

func extractRangeFilter(priceRange string) (*Range, error) {
	if priceRange == "" {
		return &Range{
			Min: -1,
			Max: -1,
		}, nil
	}
	split := strings.Split(priceRange, ",")
	if len(split) != 2 {
		return nil, errors.New("range is invalid")
	}
	var min, max float64
	var err error
	if split[0] == "" {
		min = -1
	} else {
		min, err = strconv.ParseFloat(split[0], 64)
		if err != nil {
			return nil, err
		}
	}
	if split[1] == "" {
		max = -1
	} else {
		max, err = strconv.ParseFloat(split[1], 64)
		if err != nil {
			return nil, err
		}
	}
	return &Range{
		Min: min,
		Max: max,
	}, nil
}

func extractBenefitsFilter(benefits string) (*[]string, error) {
	if benefits == "" {
		return &[]string{}, nil
	}
	b := strings.Split(benefits, ",")
	return &b, nil
}

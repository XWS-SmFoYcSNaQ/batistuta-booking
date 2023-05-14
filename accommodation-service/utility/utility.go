package utility

import "time"

func ParseISOString(isoString string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	parsedTime, err := time.Parse(layout, isoString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

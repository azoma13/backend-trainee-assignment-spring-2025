package service

import (
	"fmt"
	"time"
)

func ParseDate(date string) (*time.Time, error) {

	parsedDate, err := time.Parse(time.DateTime, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	return &parsedDate, nil
}

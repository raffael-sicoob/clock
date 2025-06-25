package utils

import (
	"time"
)

func FormatTime(currentDate string, currentTime int64) (string, time.Time) {
	parsedDate, err := time.Parse(time.RFC3339, currentDate)
	if err != nil {
		return "", time.Time{}
	}

	duration := time.Duration(currentTime) * time.Millisecond

	newDate := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local).Add(duration)
	return newDate.Format("02/01/2006 15:04:05"), newDate
}

func FormatDuration(duration int32) string {

	parsedDuration := time.Duration(duration) * time.Millisecond

	return parsedDuration.String()
}

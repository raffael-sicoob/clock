package utils

import "time"



func FormatTime(currentDate string, currentTime uint32) string {
	parsedDate, err := time.Parse(time.RFC3339, currentDate)
	if err != nil {
		return ""
	}

	duration := time.Duration(currentTime) * time.Millisecond

	newDate := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0,0,0,0, time.Local).Add(duration)
	return newDate.Format("02/01/2006 15:04:05")
}
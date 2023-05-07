package utils

import (
	"github.com/araddon/dateparse"
	"time"
)

func GetTime(strTime string) (time.Time, error) {
	t, err := dateparse.ParseAny(strTime)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func InsideInterval(date, startDate, endDate time.Time) bool {
	if date.Unix() >= startDate.Unix() && date.Unix() <= endDate.Unix() {
		return true
	}
	return false
}

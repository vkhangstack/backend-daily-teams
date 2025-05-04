package utils

import "time"

func FirstMonth() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
}
func LastMonth() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month()+1, 0, 23, 59, 0, 0, time.UTC)
}

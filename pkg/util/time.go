package util

import "time"

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTimeOfDay(d)
}

func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

func GetFirstDateOfWeek(d time.Time) time.Time {
    offset := int(time.Monday - d.Weekday())
    if offset > 0 {
        offset = -6
    }
    
    return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).AddDate(0, 0, offset)
}

func GetZeroTimeOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

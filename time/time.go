package time

import "time"

// todo: FirstLastDay(week int): 0 this week, -1 last week?
// lastWeek returns the first day and last day of last week using dateFormat
func LastWeek() (firstDay, lastDay string) {
	const dateFormat = "2006-01-02"

	now := time.Now()
	weekday := now.Weekday()

	firstDay = now.AddDate(0, 0, -7-int(weekday)).Format(dateFormat)
	lastDay = now.AddDate(0, 0, -1-int(weekday)).Format(dateFormat)
	return firstDay, lastDay
}

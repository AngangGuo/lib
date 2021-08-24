package time

import "time"

// LastWeek returns the first day and last day of last week in YYYY-MM-DD format
// todo: FirstLastDay(week int): 0 this week, -1 last week?
func LastWeek() (firstDay, lastDay string) {
	return firstLastDay(-1)
}

// firstLastDay returns the first day and last day of the week in YYYY-MM-DD format
// theWeek: 0 - this week, -1 - last week
func firstLastDay(theWeek int) (firstDay, lastDay string) {
	const dateFormat = "2006-01-02"

	now := time.Now()
	weekday := now.Weekday()

	// todo: 7*theweek + sunday - weekday
	firstDay = now.AddDate(0, 0, -7-int(weekday)).Format(dateFormat)
	// todo: 7*theweek + saturday - weekday
	lastDay = now.AddDate(0, 0, -1-int(weekday)).Format(dateFormat)
	return firstDay, lastDay
}

package workweek

import (
	"fmt"
	"time"
)

// getWorkWeek takes a date and returns the start (Monday) and end (Friday) of the week
func getWorkWeek(startDate time.Time) (start, end time.Time) {
	// Get the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	weekday := int(startDate.Weekday())

	// Calculate the difference from Monday (1)
	offset := (weekday + 6) % 7

	// Find Monday of the current week
	start = startDate.AddDate(0, 0, -offset)

	// Friday is 4 days after Monday
	end = start.AddDate(0, 0, 4)

	return start, end
}
func getNextMonday(date time.Time) time.Time {
	// Get the current weekday (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
	weekday := int(date.Weekday())

	// If today is after Monday, return the next Monday
	if weekday > 1 {
		// Add days to go to the next Monday
		daysUntilNextMonday := 8 - weekday
		return date.AddDate(0, 0, daysUntilNextMonday)
	}

	// If today is Monday or earlier, return this week's Monday
	return date
}

// assignWeeksToCandidates maps each candidate to a unique workweek starting from the current week
func GetAssignmentList(names []string) [][]string {
	currentDate := time.Now()
	startDate := getNextMonday(currentDate)
	_assignments := make([][]string, 0)
	for i, name := range names {
		// Calculate the start of the week for each candidate
		weekStart := startDate.AddDate(0, 0, i*7) // Add i weeks to the startDate
		startOfWeek, endOfWeek := getWorkWeek(weekStart)

		// Print the candidate's name along with their workweek (Monday to Friday)
		_month := startOfWeek.Month().String()
		_name := name
		_workWeek := fmt.Sprintf("%s to %s", startOfWeek.Format("02-Jan-2006"), endOfWeek.Format("02-Jan-2006"))
		_innerSlice := []string{_month, _name, _workWeek}
		_assignments = append(_assignments, _innerSlice)
	}
	return _assignments
}

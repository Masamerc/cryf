package dummy

import (
	"math/rand"
	"strings"
	"time"
)

// Generate random date string given a date range
func RandomDate(dateRange string) string {
	var start time.Time
	var rangeInSeconds int

	if dateRange == "all" {
		start = time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)
		rangeInSeconds = 60 * 60 * 24 * 365 * 20 // 20 years
	} else {
		// Split the date range into start and end
		dates := strings.Split(dateRange, "-")
		start, _ := time.Parse("2006-01-02", strings.TrimSpace(dates[0]))
		end, _ := time.Parse("2006-01-02", strings.TrimSpace(dates[1]))

		// Calculate the range in seconds
		rangeInSeconds = int(end.Sub(start).Seconds())
	}

	// Generate a random number within the range
	randNum := rand.Intn(rangeInSeconds)

	// Add the random number of seconds to the start date
	randomDate := start.Add(time.Duration(randNum) * time.Second)

	// Format the date as a string and return it
	return randomDate.Format("2006-01-02")
}

// Generate random amount based on the type of data
func RandomAmount(dataType int) int {
	switch {
	case dataType == SALES:
		return rand.Intn(100000) + 1
	default:
		return 0
	}
}

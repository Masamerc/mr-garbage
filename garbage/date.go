package garbage

import (
	"time"
)

func GetTomorrowWeekDayJst() string {
	// Load the JST timezone location
	jstLocation, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// Get the weekday tomorrow in JST
	jstTime := time.Now().In(jstLocation)
	jstTimeTomorrow := jstTime.Add(24 * time.Hour)
	jstWeekdayTomorrow := jstTimeTomorrow.Weekday().String()
	return jstWeekdayTomorrow[:3]
}

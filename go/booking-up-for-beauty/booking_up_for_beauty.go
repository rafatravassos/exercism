package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05" // formato da data de entrada
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05" // formato da data de entrada
	t, _ := time.Parse(layout, date)
	x := time.Now()
	return x.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05" // formato da data de entrada
	t, _ := time.Parse(layout, date)

	return t.Hour() >= 12 && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05" // formato da data de entrada
	t, _ := time.Parse(layout, date)
	layout = "Monday, January 2, 2006, at 15:04" // formato da data de entrada

	return fmt.Sprintf("You have an appointment on %s.", t.Format(layout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "1/2/2006" // formato da data de entrada
	currentYear := time.Now().Year()
	data := fmt.Sprintf("09/15/%d", currentYear)

	t, _ := time.Parse(layout, data)
	return t
}

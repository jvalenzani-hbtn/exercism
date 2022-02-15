package booking

import (
    "time"
    "fmt"
    )

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    const layout = "1/2/2006 15:04:05"
    tm, _ := time.Parse(layout, date)
    return tm
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    const layout = "January 2, 2006 15:04:05"
    tm, _ := time.Parse(layout, date)
	return tm.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    const layout = "Monday, January 2, 2006 15:04:05"
    tm, _ := time.Parse(layout, date)
    if(tm.Hour() >= 12 && tm.Hour() < 18){
        return true
    }
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    const layout = "1/2/2006 15:04:05"
    tm, _ := time.Parse(layout, date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", tm.Weekday(), tm.Month(), tm.Day(), tm.Year(), tm.Hour(), tm.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

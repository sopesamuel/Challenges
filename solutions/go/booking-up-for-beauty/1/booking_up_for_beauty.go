package booking

import(
    "time"
    "fmt"
) 


// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/2/2006 15:04:05", date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    result, _ := time.Parse("January 2, 2006 15:04:05", date)
    now := time.Now()

    if now.After(result){
        return true
    } else{
        return false
    }
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	
    result, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

    if result.Hour() >= 12 && result.Hour() < 18 {
        return true
    } else {
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	  t, _ := time.Parse("1/2/2006 15:04:05", date) 
    
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    t, _ := time.Parse("1/2/2006", "9/15/2020")
	return t.AddDate(time.Now().Year() - 2020 , 0, 0)
}

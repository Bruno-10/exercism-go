package booking

import (
    "strconv"
    "time"
)


// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {  
    const layout = "1/02/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    const passedLayout = "January 2, 2006 15:04:05"
    today := time.Now()
    givenDate, err := time.Parse(passedLayout, date)
    if err != nil {
        panic(err)
    }
    return givenDate.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    const iaaLayout = "Monday, January 2, 2006 15:04:05"
    givenDate, err := time.Parse(iaaLayout, date)
    if err != nil {
        panic(err)
    }

    hour := givenDate.Hour()

    return hour >= 12 && hour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    const layout = "1/2/2006 15:04:05"
    parsedDate, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    formatedDate := parsedDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
    return formatedDate
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    yearString := "15/09/" + strconv.Itoa(currentYear)
    date, err := time.Parse("02/01/2006", yearString)
    if err != nil {
        panic(err)
    }
    return date
}

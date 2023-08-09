package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
    var returnlice []Record
   for i := 0; i < len(in); i++ {
        if predicate(in[i]) {
            returnlice = append(returnlice, in[i])
        }
   }
   return returnlice
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
        return r.Day <= p.To && r.Day >= p.From
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
    return func(r Record) bool {
        return r.Category == c
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    records := Filter(in, ByDaysPeriod(p))
    var count float64
    for i := 0; i < len(records); i++ {
        count += records[i].Amount
    }
    return count
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
    recordsByDay := Filter(in, ByDaysPeriod(p))
    if len(recordsByDay) == 0 {
       return 0, nil 
    }
    recordsByCategory := Filter(recordsByDay, ByCategory(c))
    if len(recordsByCategory) == 0 {
        return 0, errors.New(fmt.Sprintf("unknown category %s", c))
    }
    var count float64
    for i := 0; i < len(recordsByCategory); i++ {
        count += recordsByCategory[i].Amount
    }
    return count, nil
}

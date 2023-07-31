package dates

import "time"

const (
	// DEFAULTDATEFORMAT is the date format used in Siena
	DEFAULTDATEFORMAT = "2006-01-02"
)

// Convert time.Time to string
// The function takes a time.Time object as input and returns a string representation of it.
func ToString(t time.Time) string {
	//fmt.Printf("t: %v\n", t)
	return toString(t)
}

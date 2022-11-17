package dates

import "time"

const (
	// DEFAULTDATEFORMAT is the date format used in Siena
	DEFAULTDATEFORMAT = "2006-01-02"
)

// Convert time.Time to string
func TimeToString(t time.Time) string {
	//fmt.Printf("t: %v\n", t)
	return t.Format(DEFAULTDATEFORMAT)
}
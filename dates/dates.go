package dates

import "time"

// Convert time.Time to string
func toString(t time.Time) string {
	//fmt.Printf("t: %v\n", t)
	return t.Format(DEFAULTDATEFORMAT)
}

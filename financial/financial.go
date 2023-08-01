package financial

import (
	"time"

	xlogger "github.com/mt1976/appFrame/logs"
)

var xlogs xlogger.XLogger

func init() {
	xlogs = xlogger.New()
}

// The function wibbleDate adjusts the input date to the next weekday if it falls on a Saturday or
// Sunday.
func wibbleDate(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

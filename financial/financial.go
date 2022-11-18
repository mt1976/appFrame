package financial

import (
	"time"
)

func wibbleDate(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

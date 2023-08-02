package financial

import (
	"time"

	xlogger "github.com/mt1976/appFrame/logs"
	xmath "github.com/mt1976/appFrame/math"
	xmock "github.com/mt1976/appFrame/mock"
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

func settlementDays(ccy1 string, ccy2 string) (int, error) {

	// Validate the two currencues using the mock package
	days1, err := xmock.GetCurrencyInfo(ccy1)
	if err != nil {
		return -1, err
	}
	days2, err := xmock.GetCurrencyInfo(ccy2)
	if err != nil {
		return -1, err
	}

	// Calculate the settlement days
	return xmath.Max(days1.SpotDays, days2.SpotDays), nil
}

func settlementDaysVia(ccy1 string, via string, ccy2 string) (int, error) {
	days1, err := settlementDays(ccy1, via)
	if err != nil {
		return -1, err
	}
	days2, err := settlementDays(via, ccy2)
	if err != nil {
		return -1, err
	}

	// Calculate the settlement days
	return xmath.Max(days1, days2), nil
}

func settlementDate(ccy1 string, ccy2 string, inDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := settlementDays(ccy1, ccy2)
	if err != nil {
		return time.Now(), err
	}

	// Adjust the date
	return wibbleDate(inDate.AddDate(0, 0, days)), nil
}

func settlementDateVia(ccy1 string, via string, ccy2 string, inDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := settlementDaysVia(ccy1, via, ccy2)
	if err != nil {
		return time.Now(), err
	}

	// Adjust the date
	return wibbleDate(inDate.AddDate(0, 0, days)), nil
}

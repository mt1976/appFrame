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

// The function adjustSettlementForWeekends adjusts the input date to the next weekday if it falls on a Saturday or
// Sunday.
func adjustSettlementForWeekends(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

func getSettlementDaysCCY(ccy1 string) (int, error) {

	// Validate the two currencues using the mock package
	days1, err := xmock.GetCurrencyInfo(ccy1)
	if err != nil {
		xlogs.WithFields(xlogger.Fields{"CCY": ccy1, "ERROR": err.Error()}).Warn("Settlement Days Issue")
		return -1, err
	}

	// Calculate the settlement days
	return days1.SpotDays, nil
}

func getSettlementDaysPAIR(ccy1 string, ccy2 string) (int, error) {

	// Validate the two currencues using the mock package
	days1, err := getSettlementDaysCCY(ccy1)
	if err != nil {
		return -1, err
	}
	days2, err := getSettlementDaysCCY(ccy2)
	if err != nil {
		return -1, err
	}

	// Calculate the settlement days
	return xmath.Max(days1, days2), nil
}

func getSettlementDaysCROSS(ccy1 string, via string, ccy2 string) (int, error) {
	days1, err := getSettlementDaysPAIR(ccy1, via)
	if err != nil {
		return -1, err
	}
	days2, err := getSettlementDaysPAIR(via, ccy2)
	if err != nil {
		return -1, err
	}

	// Calculate the settlement days
	return xmath.Max(days1, days2), nil
}

func getSettlementDateCCY(ccy1 string, inDate time.Time) (time.Time, error) {
	// Calculate the settlement days

	days, err := getSettlementDaysCCY(ccy1)
	if err != nil {
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(inDate.AddDate(0, 0, days)), nil
}

func getSettlementDatePAIR(ccy1 string, ccy2 string, inDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := getSettlementDaysPAIR(ccy1, ccy2)
	if err != nil {
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(inDate.AddDate(0, 0, days)), nil
}

func getSettlementDateCROSS(ccy1 string, via string, ccy2 string, inDate time.Time) (time.Time, error) {
	// Calculate the settlement days
	days, err := getSettlementDaysCROSS(ccy1, via, ccy2)
	if err != nil {
		return time.Now(), err
	}

	// Adjust the date
	return adjustSettlementForWeekends(inDate.AddDate(0, 0, days)), nil
}

package financial

import (
	"fmt"
	"strconv"
	"time"

	xmock "github.com/mt1976/appFrame/mock"
)

// The "FinDate" type represents a date with various properties and formats.
// @property {string} Code - The "Code" property is a string that represents a code associated with the
// date.
// @property {string} Name - The "Name" property is a string that represents the name of the date.
// @property FinDate - The "FinDate" property is of type "time.Time" and represents a specific date and time.
// @property {int} Sort - The "Sort" property is an integer that is used to determine the order in
// which the dates should be sorted. It can be used to sort the dates in ascending or descending order
// based on this value.
// @property {string} Simple - The "Simple" property is a string that represents the date in a
// simplified format. It could be a formatted string that only includes the day, month, and year of the
// date.
// @property {string} External - The "External" property in the FinDate struct is a string that represents
// an external reference or identifier related to the date.
// @property {string} Human - The "Human" property in the "FinDate" struct represents the date in a
// human-readable format. It is likely used to display the date to users in a more understandable way,
// such as "January 1, 2022".
type FinDate struct {
	Code     string
	Name     string
	Date     time.Time
	Sort     int
	Simple   string
	External string
	Human    string
}

// The function `GetDateFromTenor` calculates the settlement date based on the given tenor, trade date, and
// currency.
func GetDateFromTenor(tenor Tenor, tradeDate time.Time, ccy ...string) (time.Time, error) {

	if len(ccy) == 0 {
		return time.Now(), fmt.Errorf("no currency provided")
	}
	fmt.Printf("Tenor [%v] Trade Date [%v] Ccys %v [%v]\n", tenor.String(), tradeDate.Format("2006-01-02"), ccy, len(ccy))

	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y
	// loop thgouth currencies

	if !xmock.IsValidPeriod(tenor.String()) {
		return time.Now(), fmt.Errorf("invalid tenor [%s]", tenor.String())
	}

	spotDays := 0

	for _, c := range ccy {
		ccySpot, spotError := getSettlementDaysCCY(c)
		if spotError != nil {
			return time.Now(), spotError
		}
		if ccySpot > spotDays {
			spotDays = ccySpot
		}
	}

	switch tenor.term {
	case "SP":
		return tradeDate.AddDate(0, 0, spotDays), nil
	case "ON":
		return tradeDate.AddDate(0, 0, spotDays-1), nil
	case "TN":
		return tradeDate.AddDate(0, 0, 1), nil
	case "TD":
		return tradeDate, nil
	}

	tenorPeriod, err := tenorToDuration(tenor)
	if err != nil {
		return time.Now(), err
	}

	rtn := tradeDate.AddDate(0, 0, spotDays).Add(tenorPeriod)
	return rtn, nil
}

// The function `tenorToDuration` converts a financial tenor (e.g., "1D", "2W", "3M", "4Y") into a
// corresponding time duration.
func tenorToDuration(tenor Tenor) (time.Duration, error) {
	term := tenor.String()
	if len(term) < 2 {
		return 0, fmt.Errorf("invalid term length [%s]", term)
	}

	valueStr := term[:len(term)-1]
	unit := term[len(term)-1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("invalid term prefix [%s]", term)
	}

	switch unit {
	case 'D':
		return time.Duration(value) * 24 * time.Hour, nil
	case 'W':
		return time.Duration(value) * 7 * 24 * time.Hour, nil
	case 'M':
		return time.Duration(value) * 30 * 24 * time.Hour, nil // Assuming 30 days per month
	case 'Y':
		return time.Duration(value) * 365 * 24 * time.Hour, nil // Assuming 365 days per year
	default:
		return 0, fmt.Errorf("invalid term unit: %c", unit)
	}
}

// The function `GetLadder` takes a pivot date and a list of currency codes as input, and returns a
// list of dates based on a rate ladder.
func GetLadder(pivotDate time.Time, ccy ...string) []FinDate {
	if len(ccy) == 0 {
		return []FinDate{}
	}
	var DateList []FinDate
	rateLadder := xmock.Ladder
	xmock.RateValueToString(rateLadder)
	fmt.Printf("rateLadder: %v\n", rateLadder)
	fmt.Printf("ccy: %v\n", ccy)
	fmt.Printf("pivotDate: %v\n", pivotDate.Format("2006-01-02"))
	// range over the ladder
	for _, ladder := range rateLadder {
		thisTenor, err := NewTenor(ladder.Code)
		if err != nil {
			fmt.Printf("Error [%v]\n", err)
		}
		date, err := GetDateFromTenor(thisTenor, pivotDate, ccy...)
		if err != nil {
			fmt.Printf("Error [%v]\n", err)
		}
		fmt.Printf("thisTenor: [%v] [%v] -> [%v]\n", ladder.Code, thisTenor.String(), date.Format("2006-01-02"))
		di := FinDate{}
		di.Code = ladder.Code
		di.Name = ladder.Name
		di.Date = date
		di.Sort = ladder.Index
		di.Simple = date.Format("01/02/2006")
		di.External = date.Format("2006-01-02")
		di.Human = date.Format("Mon 02 Jan 2006")
		DateList = append(DateList, di)
	}
	return DateList
}

func getLadderCCY(pivotdate time.Time, ccy ...string) []FinDate {
	// TODO - this is a stub, please write logic
	return GetLadder(pivotdate, ccy...)
}

// The function `GetTenorFromDate` takes a date and optional currency as input and returns the
// corresponding tenor or an error.
func GetTenorFromDate(inDate time.Time, ccy ...string) (Tenor, error) {
	// TODO - TEST
	// TEST

	if len(ccy) == 0 {
		return Tenor{}, fmt.Errorf("no currency provided")
	}
	// Get Spot for the currency
	spotDate, err := GetDateFromTenor(Tenor{term: "SP"}, inDate, ccy...)
	if err != nil {
		return Tenor{}, err
	}
	// get list of tenors
	tenorList := GetLadder(spotDate, ccy...)
	// loop through tenors
	for _, tenor := range tenorList {
		if tenor.Date.Equal(inDate) {
			rtn, err := NewTenor(tenor.Code)
			if err != nil {
				return Tenor{}, err
			}
			return rtn, nil
		}
		// Now check if the tenor date is greater than the input date but less than the next tenor date
		// if so, return the tenor
		if tenor.Date.Before(inDate) {
			// get the previous tenor
			previousTenor := tenorList[tenor.Sort-1]
			rtn, err := NewTenor(previousTenor.Code)
			if err != nil {
				return Tenor{}, err
			}
			return rtn, nil
		}
	}
	return Tenor{}, nil
}

func getTenorFromDateCCY(inDate, pivotDate time.Time, ccy string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	return Tenor{}, nil
}

func getTenorFromDateCCYPAIR(inDate, pivotDate time.Time, ccy1 string, ccy2 string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	return Tenor{}, nil
}

func getTenorFromDateCCYCROSS(inDate, pivotDate time.Time, ccy1 string, via string, ccy2 string) (Tenor, error) {
	// TODO - this is a stub, please write logic
	return Tenor{}, nil
}

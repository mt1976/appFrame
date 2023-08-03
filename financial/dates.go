package financial

import (
	"fmt"
	"strconv"
	"time"
	"unicode"
)

func validateAndFormatTerm(term string) (string, error) {
	//Validates that the term string is valid
	// Validation is that the string is at least 2 characters long, and the last character is a valid unit
	// i.e. D, W, M, Y
	if len(term) < 2 {
		return "", fmt.Errorf("invalid term: %s", term)
	}
	unit := term[len(term)-1]
	unit = byte(unicode.ToUpper(rune(unit)))
	factor := term[:len(term)-1]

	_, err := strconv.Atoi(factor)
	if err != nil {
		return "", fmt.Errorf("Supplied value %s is not a number\n", factor)
	}

	clean := fmt.Sprintf("%s%c", factor, unit)

	switch unit {
	case 'D':
		return clean, nil
	case 'W':
		return clean, nil
	case 'M':
		return clean, nil
	case 'Y':
		return clean, nil
	default:
		return "", fmt.Errorf("invalid term unit: %c", unit)
	}
}

func getTenorDateCCY(inTerm string, ccy string) (time.Time, error) {
	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y

	inTerm, termError := validateAndFormatTerm(inTerm)
	if termError != nil {
		return time.Now(), termError
	}

	spotDate, spotError := getSettlementDateCCY(ccy, time.Now())
	if spotError != nil {
		return time.Now(), spotError
	}
	days, err := bankingTermToDuration(inTerm)
	if err != nil {
		return time.Now(), err
	}
	return spotDate.Add(days), nil
}

func getTenorDateCCYPAIR(inTerm string, ccy1 string, ccy2 string) (time.Time, error) {
	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y
	return time.Now(), nil
}

func bankingTermToDuration(term string) (time.Duration, error) {
	if len(term) < 2 {
		return 0, fmt.Errorf("invalid term: %s", term)
	}

	valueStr := term[:len(term)-1]
	unit := term[len(term)-1]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("invalid term: %s", term)
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

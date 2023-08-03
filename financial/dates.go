package financial

import "time"

func getTenorDateCCY(inTerm string, ccy string) time.Time {
	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y
	return time.Now()
}

func getTenorDateCCYPAIR(inTerm string, ccy1 string, ccy2 string) time.Time {
	// Calculate the settlement days, and adjust the date based on the term string provided i.e. 1D, 1W, 1M, 1Y
	return time.Now()
}

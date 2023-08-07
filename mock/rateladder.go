package mock

import "fmt"

type Rung struct {
	Code        string
	Name        string
	Alternative string
	Index       int
}

var Ladder map[string]Rung

func init() {
	Ladder = make(map[string]Rung)
	Ladder["ON"] = Rung{Code: "ON", Name: "Overnight", Index: 1}
	Ladder["TD"] = Rung{Code: "TD", Name: "Today", Index: 2}
	Ladder["TN"] = Rung{Code: "TN", Name: "Tom/Next", Index: 3}
	Ladder["SP"] = Rung{Code: "SP", Name: "Spot", Index: 4}
	Ladder["1W"] = Rung{Code: "1W", Name: "1 Week", Alternative: "7D", Index: 5}
	Ladder["2W"] = Rung{Code: "2W", Name: "2 Weeks", Index: 6}
	Ladder["3W"] = Rung{Code: "3W", Name: "3 Weeks", Index: 7}
	Ladder["1M"] = Rung{Code: "1M", Name: "1 Month", Alternative: "30D", Index: 8}
	Ladder["2M"] = Rung{Code: "2M", Name: "2 Months", Index: 9}
	Ladder["3M"] = Rung{Code: "3M", Name: "3 Months", Alternative: "90D", Index: 10}
	Ladder["4M"] = Rung{Code: "4M", Name: "4 Months", Index: 11}
	Ladder["5M"] = Rung{Code: "5M", Name: "5 Months", Index: 12}
	Ladder["6M"] = Rung{Code: "6M", Name: "6 Months", Alternative: "180D", Index: 13}
	Ladder["7M"] = Rung{Code: "7M", Name: "7 Months", Index: 14}
	Ladder["8M"] = Rung{Code: "8M", Name: "8 Months", Index: 15}
	Ladder["9M"] = Rung{Code: "9M", Name: "9 Months", Index: 16}
	Ladder["10M"] = Rung{Code: "10M", Name: "10 Months", Index: 17}
	Ladder["11M"] = Rung{Code: "11M", Name: "11 Months", Index: 18}
	Ladder["1Y"] = Rung{Code: "1Y", Name: "1 Year", Alternative: "12M", Index: 19}
	Ladder["15M"] = Rung{Code: "15M", Name: "13 Months", Index: 20}
	Ladder["18M"] = Rung{Code: "18M", Name: "18 Months", Index: 21}
	Ladder["21M"] = Rung{Code: "21M", Name: "21 Months", Index: 22}
	Ladder["2Y"] = Rung{Code: "2Y", Name: "2 Years", Index: 23}
	Ladder["3Y"] = Rung{Code: "3Y", Name: "3 Years", Index: 24}
	Ladder["4Y"] = Rung{Code: "4Y", Name: "4 Years", Index: 25}
	Ladder["5Y"] = Rung{Code: "5Y", Name: "5 Years", Index: 26}
	Ladder["6Y"] = Rung{Code: "6Y", Name: "6 Years", Index: 27}
	Ladder["7Y"] = Rung{Code: "7Y", Name: "7 Years", Index: 28}
	Ladder["8Y"] = Rung{Code: "8Y", Name: "8 Years", Index: 29}
	Ladder["9Y"] = Rung{Code: "9Y", Name: "9 Years", Index: 30}
	Ladder["10Y"] = Rung{Code: "10Y", Name: "10 Years", Index: 31}
}

func GetRateLadderList() []string {
	rtn := []string{}
	for k := range Ladder {
		rtn = append(rtn, k)
	}
	//rtn.sort()
	return rtn
}

func IsValidPeriod(in string) bool {
	_, ok := Ladder[in]
	return ok
}

func GetRateLadderByIndex(index int) Rung {
	for _, v := range Ladder {
		if v.Index == index {
			return v
		}
	}
	return Rung{}
}

func test() bool {
	noitems := len(Ladder)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		fmt.Printf("rate ladder info: %v\n", rli)
	}
	return true
}

func RateValueToString(R map[string]Rung) string {
	output := ""
	//	noItems := len(R)
	noitems := len(Ladder)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		// add to output
		output += fmt.Sprintf("%v:%v,", rli.Code, rli.Index)
	}

	return output
}

func GetTenorInfo(tenor string) (Rung, error) {
	return Ladder[tenor], nil
}

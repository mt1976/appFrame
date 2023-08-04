package mock

import "fmt"

type RateLadderInfo struct {
	Code        string
	Name        string
	Alternative string
	Index       int
}

var RateLadderInfoMap map[string]RateLadderInfo

func init() {
	RateLadderInfoMap = make(map[string]RateLadderInfo)
	RateLadderInfoMap["ON"] = RateLadderInfo{Code: "ON", Name: "Overnight", Index: 1}
	RateLadderInfoMap["TD"] = RateLadderInfo{Code: "TD", Name: "Today", Index: 2}
	RateLadderInfoMap["TN"] = RateLadderInfo{Code: "TN", Name: "Tom/Next", Index: 3}
	RateLadderInfoMap["SP"] = RateLadderInfo{Code: "SP", Name: "Spot", Index: 4}
	RateLadderInfoMap["1W"] = RateLadderInfo{Code: "1W", Name: "1 Week", Alternative: "7D", Index: 5}
	RateLadderInfoMap["2W"] = RateLadderInfo{Code: "2W", Name: "2 Weeks", Index: 6}
	RateLadderInfoMap["3W"] = RateLadderInfo{Code: "3W", Name: "3 Weeks", Index: 7}
	RateLadderInfoMap["1M"] = RateLadderInfo{Code: "1M", Name: "1 Month", Alternative: "30D", Index: 8}
	RateLadderInfoMap["2M"] = RateLadderInfo{Code: "2M", Name: "2 Months", Index: 9}
	RateLadderInfoMap["3M"] = RateLadderInfo{Code: "3M", Name: "3 Months", Alternative: "90D", Index: 10}
	RateLadderInfoMap["4M"] = RateLadderInfo{Code: "4M", Name: "4 Months", Index: 11}
	RateLadderInfoMap["5M"] = RateLadderInfo{Code: "5M", Name: "5 Months", Index: 12}
	RateLadderInfoMap["6M"] = RateLadderInfo{Code: "6M", Name: "6 Months", Alternative: "180D", Index: 13}
	RateLadderInfoMap["7M"] = RateLadderInfo{Code: "7M", Name: "7 Months", Index: 14}
	RateLadderInfoMap["8M"] = RateLadderInfo{Code: "8M", Name: "8 Months", Index: 15}
	RateLadderInfoMap["9M"] = RateLadderInfo{Code: "9M", Name: "9 Months", Index: 16}
	RateLadderInfoMap["10M"] = RateLadderInfo{Code: "10M", Name: "10 Months", Index: 17}
	RateLadderInfoMap["11M"] = RateLadderInfo{Code: "11M", Name: "11 Months", Index: 18}
	RateLadderInfoMap["1Y"] = RateLadderInfo{Code: "1Y", Name: "1 Year", Alternative: "12M", Index: 19}
	RateLadderInfoMap["15M"] = RateLadderInfo{Code: "15M", Name: "13 Months", Index: 20}
	RateLadderInfoMap["18M"] = RateLadderInfo{Code: "18M", Name: "18 Months", Index: 21}
	RateLadderInfoMap["21M"] = RateLadderInfo{Code: "21M", Name: "21 Months", Index: 22}
	RateLadderInfoMap["2Y"] = RateLadderInfo{Code: "2Y", Name: "2 Years", Index: 23}
	RateLadderInfoMap["3Y"] = RateLadderInfo{Code: "3Y", Name: "3 Years", Index: 24}
	RateLadderInfoMap["4Y"] = RateLadderInfo{Code: "4Y", Name: "4 Years", Index: 25}
	RateLadderInfoMap["5Y"] = RateLadderInfo{Code: "5Y", Name: "5 Years", Index: 26}
	RateLadderInfoMap["6Y"] = RateLadderInfo{Code: "6Y", Name: "6 Years", Index: 27}
	RateLadderInfoMap["7Y"] = RateLadderInfo{Code: "7Y", Name: "7 Years", Index: 28}
	RateLadderInfoMap["8Y"] = RateLadderInfo{Code: "8Y", Name: "8 Years", Index: 29}
	RateLadderInfoMap["9Y"] = RateLadderInfo{Code: "9Y", Name: "9 Years", Index: 30}
	RateLadderInfoMap["10Y"] = RateLadderInfo{Code: "10Y", Name: "10 Years", Index: 31}
}

func GetRateLadderList() []string {
	rtn := []string{}
	for k := range RateLadderInfoMap {
		rtn = append(rtn, k)
	}
	//rtn.sort()
	return rtn
}

func IsValidPeriod(in string) bool {
	_, ok := RateLadderInfoMap[in]
	return ok
}

func GetRateLadderByIndex(index int) RateLadderInfo {
	for _, v := range RateLadderInfoMap {
		if v.Index == index {
			return v
		}
	}
	return RateLadderInfo{}
}

func test() bool {
	noitems := len(RateLadderInfoMap)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		fmt.Printf("rate ladder info: %v\n", rli)
	}
	return true
}

func RateValueToString(R map[string]RateLadderInfo) string {
	output := ""
	//	noItems := len(R)
	noitems := len(RateLadderInfoMap)
	for i := 1; i <= noitems; i++ {
		rli := GetRateLadderByIndex(i)
		// add to output
		output += fmt.Sprintf("%v:%v,", rli.Code, rli.Index)
	}

	return output
}

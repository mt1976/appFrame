package mock

import (
	"fmt"
	"log"
	"time"
)

// CurrencyInfo represents information about the IBAN format for a specific country.
type CurrencyInfo struct {
	Code               string // The ISO currency code
	SpotDays           int    // The number of spot days for the currency
	Name               string // The name of the currency
	Character          string // The character of the currency
	DPS                int    // The number of decimal places for the currency
	QuoteDPS           int    // The number of decimal places for the currency when quoting
	Type               string // The type of currency
	MajorUnit          string // The major unit of the currency
	MinorUnit          string // The minor unit of the currency
	MinorCharacter     string // The minor character of the currency
	ISONumericCode     string // The ISO numeric code of the currency
	KnownAs            string // The known as name of the currency
	YearOfIntroduction int    // The year the currency was introduced
}

const (
	Currency = "Currency"
	Crypto   = "Crypto"
	Metals   = "Psuedo"
	Testing  = "Testing"
)

var CurrencyInfoMap map[string]CurrencyInfo

func init() {

	//fmt.Println("mock init")
	CurrencyInfoMap = make(map[string]CurrencyInfo)
	CurrencyInfoMap["EUR"] = CurrencyInfo{Code: "EUR", SpotDays: 2, Name: "Euro", Character: "€", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Euro", MinorUnit: "Cent", ISONumericCode: "978", KnownAs: "Euros", MinorCharacter: "c", YearOfIntroduction: 1999}
	CurrencyInfoMap["USD"] = CurrencyInfo{Code: "USD", SpotDays: 1, Name: "US Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "840", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1792}
	CurrencyInfoMap["GBP"] = CurrencyInfo{Code: "GBP", SpotDays: 2, Name: "Pound Sterling", Character: "£", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Pound", MinorUnit: "Pence", ISONumericCode: "826", KnownAs: "Quids", MinorCharacter: "p", YearOfIntroduction: 800}
	CurrencyInfoMap["ZAR"] = CurrencyInfo{Code: "ZAR", SpotDays: 2, Name: "South African Rand", Character: "R", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Rand", MinorUnit: "Cent", ISONumericCode: "710", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1961}
	CurrencyInfoMap["MXN"] = CurrencyInfo{Code: "MXN", SpotDays: 3, Name: "Mexican Peso", Character: "$", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Peso", MinorUnit: "Centavo", ISONumericCode: "484", KnownAs: "Pesos", MinorCharacter: "c", YearOfIntroduction: 1993}
	CurrencyInfoMap["CAD"] = CurrencyInfo{Code: "CAD", SpotDays: 1, Name: "Canadian Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "124", KnownAs: "Loonies", MinorCharacter: "c", YearOfIntroduction: 1858}
	CurrencyInfoMap["JPY"] = CurrencyInfo{Code: "JPY", SpotDays: 2, Name: "Japanese Yen", Character: "¥", DPS: 0, QuoteDPS: 2, Type: Currency, MajorUnit: "Yen", MinorUnit: "Sen", ISONumericCode: "392", KnownAs: "Yen", MinorCharacter: "s", YearOfIntroduction: 1871}
	CurrencyInfoMap["CHF"] = CurrencyInfo{Code: "CHF", SpotDays: 2, Name: "Swiss Franc", Character: "Fr", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Franc", MinorUnit: "Rappen", ISONumericCode: "756", KnownAs: "Swissies", MinorCharacter: "rp", YearOfIntroduction: 1850}
	CurrencyInfoMap["AUD"] = CurrencyInfo{Code: "AUD", SpotDays: 2, Name: "Australian Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "036", KnownAs: "Aussie", MinorCharacter: "c", YearOfIntroduction: 1966}
	CurrencyInfoMap["INR"] = CurrencyInfo{Code: "INR", SpotDays: 2, Name: "Indian Rupee", Character: "₹", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Rupee", MinorUnit: "Paisa", ISONumericCode: "356", KnownAs: "Rupayya", MinorCharacter: "p", YearOfIntroduction: 1540}
	CurrencyInfoMap["CLF"] = CurrencyInfo{Code: "CLF", SpotDays: 2, Name: "Chilean Unidad de Fomento", Character: "UF", DPS: 4, QuoteDPS: 4, Type: Currency, MajorUnit: "Unidad de Fomento", MinorUnit: "Peso", ISONumericCode: "990", YearOfIntroduction: 1967}
	CurrencyInfoMap["CNY"] = CurrencyInfo{Code: "CNY", SpotDays: 2, Name: "Chinese Yuan Renminbi", Character: "¥", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Yuan", MinorUnit: "Fen", ISONumericCode: "156", KnownAs: "Yuan", MinorCharacter: "f", YearOfIntroduction: 1949}
	CurrencyInfoMap["IQD"] = CurrencyInfo{Code: "IQD", SpotDays: 2, Name: "Iraqi Dinar", Character: "ع.د", DPS: 3, QuoteDPS: 4, Type: Currency, MajorUnit: "Dinar", MinorUnit: "Fils", ISONumericCode: "368", KnownAs: "Dinar", MinorCharacter: "f", YearOfIntroduction: 1932}
	CurrencyInfoMap["XAG"] = CurrencyInfo{Code: "XAG", SpotDays: 2, Name: "Silver", Character: "Ag", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "961", KnownAs: "Silver"}
	CurrencyInfoMap["XAU"] = CurrencyInfo{Code: "XAU", SpotDays: 2, Name: "Gold", Character: "Au", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "959", KnownAs: "Gold"}
	CurrencyInfoMap["XTS"] = CurrencyInfo{Code: "XTS", SpotDays: 2, Name: "Testing Currency Code", Character: "¤", DPS: 4, QuoteDPS: 4, Type: Testing, MajorUnit: "Unit", MinorUnit: "Unit", ISONumericCode: "999", KnownAs: "Testing Currency Code", MinorCharacter: "u", YearOfIntroduction: 1970}
	CurrencyInfoMap["BTC"] = CurrencyInfo{Code: "BTC", SpotDays: 0, Name: "Bitcoin", Character: "₿", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Bitcoin", MinorUnit: "Satoshi", ISONumericCode: "1001", KnownAs: "Bitcoin", MinorCharacter: "s", YearOfIntroduction: 2009}
	CurrencyInfoMap["ETH"] = CurrencyInfo{Code: "ETH", SpotDays: 0, Name: "Ethereum", Character: "Ξ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Ether", MinorUnit: "Wei", ISONumericCode: "1002", KnownAs: "Ether", MinorCharacter: "w", YearOfIntroduction: 2015}
	CurrencyInfoMap["LTC"] = CurrencyInfo{Code: "LTC", SpotDays: 0, Name: "Litecoin", Character: "Ł", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Litecoin", MinorUnit: "Litetoshi", ISONumericCode: "1003", KnownAs: "Litecoin", MinorCharacter: "l", YearOfIntroduction: 2011}
	CurrencyInfoMap["XRP"] = CurrencyInfo{Code: "XRP", SpotDays: 0, Name: "Ripple", Character: "Ʀ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Ripple", MinorUnit: "Drops", ISONumericCode: "1004", KnownAs: "XRup", MinorCharacter: "d", YearOfIntroduction: 2012}
	CurrencyInfoMap["DASH"] = CurrencyInfo{Code: "DASH", SpotDays: 0, Name: "Dash", Character: "Đ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dash", MinorUnit: "Duffs", ISONumericCode: "1005", KnownAs: "Dash", MinorCharacter: "d", YearOfIntroduction: 2014}
	CurrencyInfoMap["DOGE"] = CurrencyInfo{Code: "DOGE", SpotDays: 0, Name: "Dogecoin", Character: "Ð", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dogecoin", MinorUnit: "Shibes", ISONumericCode: "1006", KnownAs: "Much Wow", MinorCharacter: "s", YearOfIntroduction: 2013}
}

func GetCurrencyInfo(code string) (CurrencyInfo, error) {

	rtn := CurrencyInfo{}
	if len(code) > 1 && len(code) <= 4 {
		rtn = CurrencyInfoMap[code]
	}

	if len(rtn.Code) == 0 {
		log.Printf("[WARN] Invalid currency code: [%s]", code)
		return CurrencyInfo{}, fmt.Errorf("invalid currency code: [%s]", code)
	}
	//fmt.Printf("CurrencyInfo: %v Age %v years\n", rtn, rtn.Age())
	return rtn, nil
}

func (C *CurrencyInfo) Age() int {
	return time.Now().Year() - C.YearOfIntroduction
	//return C.SpotDays
}

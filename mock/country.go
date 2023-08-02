package mock

import (
	"fmt"
	"log"
)

// CountryInfo represents information about the IBAN format for a specific country.
type CountryInfo struct {
	IBANLength int    // The expected length of the IBAN for the country.
	Currency   string // The currency code for the country.
	ISOCode    string // The ISO country code
	ISOCode3   string // The ISO country code
}

var CountryInfoMap map[string]CountryInfo

func init() {
	fmt.Println("mock init")
	CountryInfoMap = make(map[string]CountryInfo)
	CountryInfoMap["DE"] = CountryInfo{IBANLength: 22, Currency: "EUR", ISOCode: "DE", ISOCode3: "DEU"}
	CountryInfoMap["US"] = CountryInfo{IBANLength: 18, Currency: "USD", ISOCode: "US", ISOCode3: "USA"}
	CountryInfoMap["GB"] = CountryInfo{IBANLength: 22, Currency: "GBP", ISOCode: "GB", ISOCode3: "GBR"}
	CountryInfoMap["FR"] = CountryInfo{IBANLength: 27, Currency: "EUR", ISOCode: "FR", ISOCode3: "FRA"}
	CountryInfoMap["ES"] = CountryInfo{IBANLength: 24, Currency: "EUR", ISOCode: "ES", ISOCode3: "ESP"}
	CountryInfoMap["IT"] = CountryInfo{IBANLength: 27, Currency: "EUR", ISOCode: "IT", ISOCode3: "ITA"}
	CountryInfoMap["NL"] = CountryInfo{IBANLength: 18, Currency: "EUR", ISOCode: "NL", ISOCode3: "NLD"}
	CountryInfoMap["ZA"] = CountryInfo{IBANLength: 20, Currency: "ZAR", ISOCode: "ZA", ISOCode3: "ZAF"}
}

func GetCountryInfo(countryCode string) (CountryInfo, error) {

	rtn := CountryInfo{}
	if len(countryCode) == 2 {
		rtn = CountryInfoMap[countryCode]
	}

	if len(countryCode) == 3 {
		for _, v := range CountryInfoMap {
			if v.ISOCode3 == countryCode {
				rtn = v
			}
		}
	}

	if rtn.IBANLength == 0 {
		log.Printf("[WARN] Invalid country code: [%s]", countryCode)
		return CountryInfo{}, fmt.Errorf("invalid country code: [%s]", countryCode)
	}
	return rtn, nil
}

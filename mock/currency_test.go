package mock

import (
	"reflect"
	"testing"
)

func TestGetCurrencyInfo(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    CurrencyInfo
		wantErr bool
	}{
		{"Valid Currency Code", args{"EUR"}, CurrencyInfo{Code: "EUR", SpotDays: 2, Name: "Euro", Character: "€", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Euro", MinorUnit: "Cent", ISONumericCode: "978", KnownAs: "Euros", MinorCharacter: "c", YearOfIntroduction: 1999}, false},
		{"Invalid Currency Code", args{"XX"}, CurrencyInfo{}, true},
		{"Valid Currency Code", args{"USD"}, CurrencyInfo{Code: "USD", SpotDays: 1, Name: "US Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Currency, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "840", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1792}, false},
		{"Crypto DASH", args{"DASH"}, CurrencyInfo{Code: "DASH", SpotDays: 0, Name: "Dash", Character: "Đ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dash", MinorUnit: "Duffs", ISONumericCode: "1005", KnownAs: "Dash", MinorCharacter: "d", YearOfIntroduction: 2014}, false},
		{"Metal XAU", args{"XAU"}, CurrencyInfo{Code: "XAU", SpotDays: 2, Name: "Gold", Character: "Au", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "959", KnownAs: "Gold"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrencyInfo(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrencyInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrencyInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

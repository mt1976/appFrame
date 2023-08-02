package mock

import (
	"reflect"
	"testing"
)

func TestGetCountryInfo(t *testing.T) {
	type args struct {
		countryCode string
	}
	tests := []struct {
		name    string
		args    args
		want    CountryInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Valid Country Code", args{"DE"}, CountryInfo{IBANLength: 22, Currency: "EUR", ISOCode: "DE", ISOCode3: "DEU"}, false},
		{"Invalid Country Code", args{"XX"}, CountryInfo{}, true},
		{"Valid Country Code", args{"DEU"}, CountryInfo{IBANLength: 22, Currency: "EUR", ISOCode: "DE", ISOCode3: "DEU"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCountryInfo(tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCountryInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCountryInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

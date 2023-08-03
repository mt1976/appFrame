package financial

import (
	"reflect"
	"testing"
	"time"
)

func Test_validateAndFormatTerm(t *testing.T) {
	type args struct {
		term string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SP", args{"SP"}, "SP", false},
		{"td", args{"td"}, "TD", false},
		{"1D", args{"1D"}, "1D", false},
		{"1W", args{"1W"}, "1W", false},
		{"1M", args{"1M"}, "1M", false},
		{"1Y", args{"1Y"}, "1Y", false},
		{"1d", args{"1d"}, "1D", false},
		{"10w", args{"10w"}, "10W", false},
		{"10m", args{"10m"}, "10M", false},
		{"10X", args{"10X"}, "", true},
		{"BUMBUM", args{"BUMBUM"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAndFormatTenor(tt.args.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAndFormatTerm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateAndFormatTerm() = %v, want %v, %e", got, tt.want, err)
			}
			t.Logf("validateAndFormatTerm() = %v, want %v, error %v", got, tt.want, err)
		})
	}
}

func Test_getTenorDateCCY(t *testing.T) {
	type args struct {
		tenor     Tenor
		tradeDate time.Time
		ccy       string
	}
	SP, _ := NewTenor("SP")
	TD, _ := NewTenor("TD")

	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.

		{"SPUSD", args{SP, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), "USD"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"SPEUR", args{SP, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), "EUR"}, time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), false},
		{"TDEUR", args{TD, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), "EUR"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTenorDateCCY(tt.args.tenor, tt.args.tradeDate, tt.args.ccy)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTenorDateCCY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTenorDateCCY() = %v, want %v", got, tt.want)
			}
			t.Logf("getTenorDateCCY() = %v, want %v", got, tt.want)
		})
	}
}

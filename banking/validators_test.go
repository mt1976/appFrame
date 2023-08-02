package banking

import (
	"testing"
)

func Test_isValidIBAN(t *testing.T) {
	tests := []struct {
		iban     string
		expected bool
	}{
		// TODO: Add test cases.
		{"DE89370400440532013000", true},
		{"DE12345678901234567890", false},
		{"poo", false},
	}
	for _, tt := range tests {
		t.Run(tt.iban, func(t *testing.T) {
			actual := isValidIBAN(tt.iban)
			if actual != tt.expected {
				t.Errorf("isValidIBAN(%s): expected %v, actual %v", tt.iban, tt.expected, actual)
			}
		})
	}
}

func Test_isValidLEI(t *testing.T) {
	type args struct {
		lei string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Valid LEI", args{"529900T8BM49AURSDO55"}, true},
		{"Invalid LEI", args{"12345678901234567890"}, false},
		{"EUROBASE", args{"8755008BT1IYMFZXV751"}, true},
		{"GLIEF", args{"506700GE1G29325QX363"}, true},
		{"Bank of Ireland", args{"6354002WOGLFYPOC1W29"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidLEI(tt.args.lei); got != tt.want {
				t.Errorf("isValidLEI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidISIN(t *testing.T) {
	type args struct {
		isin string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"APPLE INC", args{"US0378331005"}, true},
		{"Invalid ISIN", args{"US0378331001"}, false},
		{"WALMART", args{"US9311421039"}, true},
		{"TEST", args{"US0378331005"}, true},
		{"BAE Systems", args{"GB0002634946"}, true},
		{"Bank of Ireland", args{"6354002WOGLFYPOC1W29"}, false},
		{"Invalid Country Code", args{"XX0002634946"}, false},
	}
	for _, tt := range tests {
		var isin ISIN
		isin.Set(tt.args.isin)
		//isin := tt.args

		t.Run(tt.name, func(t *testing.T) {
			if got := isin.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

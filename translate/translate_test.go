package translate

import (
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test()
		})
	}
}

func Test_outFormat(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"SP", args{"SP"}, "SP"},
		{"td", args{"td"}, "td"},
		//{"Space", args{"{{space}}"}, " "},
		{"Eq", args{"{{eq}}"}, "="},
		{"Gt", args{"{{gt}}"}, ">"},
		{"Null", args{"null"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outFormat(tt.args.in); got != tt.want {
				t.Errorf("outFormat() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}

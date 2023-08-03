package banking

import (
	"reflect"
	"testing"
)

func TestNewLEI(t *testing.T) {
	type args struct {
		lei string
	}
	tests := []struct {
		name    string
		args    args
		want    LEI
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Valid LEI", args{"213800A8Y1XKQMG8S714"}, LEI{"213800A8Y1XKQMG8S714"}, false},
		{"Invalid LEI", args{"213800A8Y1XKQMG8S713"}, LEI{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLEI(tt.args.lei)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLEI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLEI() = %v, want %v", got, tt.want)
			}
		})
	}
}

package temp

import (
	"reflect"
	"testing"

	xdl "github.com/mt1976/appFrame/dataloader"
)

func TestFetch(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    TempData
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", args{"test"}, TempData{"test", PathSeparator + "temp" + PathSeparator + "test", xdl.Payload{}}, false},
		{"noname", args{""}, TempData{"", "", xdl.Payload{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fetch(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() : got %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

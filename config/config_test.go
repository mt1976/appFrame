package config

import (
	"reflect"
	"testing"

	xdl "github.com/mt1976/appFrame/dataloader"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
		path string
	}
	tests := []struct {
		name string
		args args
		want xdl.Payload
	}{
		// TODO: Add test cases.
		{"test1", args{"test", ""}, *xdl.New("test", "properties", "")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("New(%v,%v) = ??, want %v", tt.args.name, tt.args.path, tt.want)

			got := New(tt.args.name, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			t.Logf("New(%v,%v) = %v, want %v", tt.args.name, tt.args.path, got, tt.want)
			Debug(got)
		})
	}
}

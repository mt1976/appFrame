package system

import "testing"

func Test_getNetworkName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"test", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNetworkName(); got != tt.want {
				t.Errorf("getNetworkName() = %v, want %v", got, tt.want)
			}
		})
	}
}

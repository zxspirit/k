package build

import "testing"

func TestServe(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestServe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Serve()
		})
	}
}

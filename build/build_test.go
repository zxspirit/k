package build

import (
	"reflect"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{"TestNewServer", args{}, &Server{
			Addr:         "localhost",
			Port:         8080,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Timeout:      10 * time.Second,
		}},
		{"TestNewServer1", args{[]Option{WithAddr("localhost"),
			WithTimeout(30 * time.Second),
			WithWriteTimeout(30 * time.Second),
			WithReadTimeout(30 * time.Second)}},
			&Server{
				Addr:         "localhost",
				Port:         8080,
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 30 * time.Second,
				Timeout:      30 * time.Second,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

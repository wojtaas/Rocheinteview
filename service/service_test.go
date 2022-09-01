package service

import (
	"rocheinteview"
	"strings"
	"testing"
)

func TestService_Ping(t *testing.T) {
	tests := []struct {
		name       string
		message    string
		wantOutput string
	}{{
		name:       "happy path",
		message:    "hello",
		wantOutput: "hello",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				config:  rocheinteview.Config{},
				Version: "1.0.2",
			}
			output := s.Ping(tt.message)
			if !strings.EqualFold(output.Echo, tt.wantOutput) {
				t.Errorf("want %s got %s", tt.wantOutput, output.Echo)
			}
		})
	}
}

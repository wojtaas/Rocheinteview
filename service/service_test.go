package service

import (
	"rocheinteview"
	"strings"
	"testing"
	"time"
)

func TestService_Ping(t *testing.T) {
	type args struct {
		version string
		message string
		env     string
	}

	tests := []struct {
		name       string
		args       args
		wantOutput rocheinteview.PingResponse
	}{
		{
			name: "happy path #1",
			args: args{
				version: "1.0.0",
				message: "hello",
				env:     "OS",
			},
			wantOutput: rocheinteview.PingResponse{
				Echo:      "hello",
				Timestamp: time.Now().Unix(),
				Env:       "OS",
				Version:   "1.0.0",
			},
		},
		{
			name: "happy path #2",
			args: args{
				version: "0.0.1",
				message: "Hey Hi Hello Some More Text",
			},
			wantOutput: rocheinteview.PingResponse{
				Echo:      "Hey Hi Hello Some More Text",
				Timestamp: time.Now().Unix(),
				Version:   "0.0.1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				config:  rocheinteview.Config{Env: tt.args.env},
				Version: tt.args.version,
			}
			output := s.Ping(tt.args.message)
			if !strings.EqualFold(output.Echo, tt.wantOutput.Echo) {
				t.Errorf("ECHO: want %s got %s", tt.wantOutput.Echo, output.Echo)
			}
			if !strings.EqualFold(output.Version, tt.wantOutput.Version) {
				t.Errorf("VERSION: want %s got %s", tt.wantOutput.Version, output.Version)
			}
			if !strings.EqualFold(output.Env, tt.wantOutput.Env) {
				t.Errorf("ENV: want %s got %s", tt.wantOutput.Env, output.Env)
			}
			if output.Timestamp != tt.wantOutput.Timestamp {
				t.Errorf("TIMESTAMP: want %d got %d", tt.wantOutput.Timestamp, output.Timestamp)
			}

		})
	}
}

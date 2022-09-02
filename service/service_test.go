package service

import (
	"github.com/stretchr/testify/assert"
	"rocheinteview"
	"testing"
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
				Env:     "OS",
				Version: "1.0.0",
			},
		},
		{
			name: "happy path #2",
			args: args{
				version: "0.0.1",
				message: "Hey Hi Hello Some More Text",
			},
			wantOutput: rocheinteview.PingResponse{
				Version: "0.0.1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				config:  rocheinteview.Config{Env: tt.args.env},
				Version: tt.args.version,
			}
			output, err := s.Ping(tt.args.message)

			assert.NoError(t, err)
			assert.NotEmpty(t, output.Echo)
			assert.Equal(t, output.Env, tt.wantOutput.Env)
			assert.Equal(t, output.Version, tt.wantOutput.Version)
		})
	}
}

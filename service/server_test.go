package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rocheinteview"
	. "rocheinteview/grpc/pb-go"
	"rocheinteview/service/mocks"
	"testing"
	"time"
)

func TestServer_Ping(t *testing.T) {
	tests := []struct {
		name       string
		service    Service
		message    string
		wantOutput *PingResponse
	}{

		{
			name:    "ok",
			message: "hello",
			wantOutput: &PingResponse{
				Echo:      "hello",
				Timestamp: time.Now().Unix(),
				Env:       "someEnv",
				Version:   "1.0.0",
			},
			service: func() *mocks.Service {
				m := new(mocks.Service)
				m.On("Ping", "hello").Return(rocheinteview.PingResponse{
					Echo:      "hello",
					Timestamp: time.Now().Unix(),
					Env:       "someEnv",
					Version:   "1.0.0",
				})
				return m
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server{service: tt.service}

			response, err := s.Ping(context.Background(), &PingRequest{Message: tt.message})
			if !reflect.DeepEqual(response, tt.wantOutput) {
				t.Errorf("s.Ping{} got= %v, want %v", response, tt.wantOutput)
			}
			assert.NoError(t, err)
		})
	}
}

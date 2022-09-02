package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"rocheinteview"
	"time"
)

type service struct {
	config  rocheinteview.Config
	Version string
}

// NewService function is used in main, to create new instance of Service
func NewService(cfg rocheinteview.Config, version string) *service {
	return &service{
		config:  cfg,
		Version: version,
	}
}

// Ping function is used by API to return echo message with additionally parameters.
func (s *service) Ping(message string) (rocheinteview.PingResponse, error) {
	postmanEchoResponse, err := http.Get(fmt.Sprintf(rocheinteview.PostmanEchoURL, message))
	if err != nil {
		return rocheinteview.PingResponse{}, errors.New(fmt.Sprintf("error occured while sending get request: %d", err))
	}

	postmanEchoBody, err := io.ReadAll(postmanEchoResponse.Body)
	if err != nil {
		return rocheinteview.PingResponse{}, errors.New(fmt.Sprintf("io.ReadAll(): %d", err))
	}

	return rocheinteview.PingResponse{
		Echo:      string(postmanEchoBody),
		Timestamp: time.Now().Unix(),
		Env:       s.config.Env,
		Version:   s.Version,
	}, nil
}

package service

import (
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
func (s *service) Ping(message string) rocheinteview.PingResponse {
	return rocheinteview.PingResponse{
		Echo:      message,
		Timestamp: time.Now().Unix(),
		Env:       s.config.Env,
		Version:   s.Version,
	}
}

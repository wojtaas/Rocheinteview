package service

import (
	"rocheinteview"
	"time"
)

type service struct {
	config  rocheinteview.Config
	Version string
}

func NewService(cfg rocheinteview.Config, version string) *service {
	return &service{
		config:  cfg,
		Version: version,
	}
}

func (s *service) Ping(message string) rocheinteview.PingResponse {
	return rocheinteview.PingResponse{
		Echo:      message,
		Timestamp: time.Now().Unix(),
		Env:       s.config.Env,
		Version:   s.Version,
	}
}

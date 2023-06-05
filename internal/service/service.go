package service

import "github.com/go-nunu/nunu-layout-base/pkg/log"

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

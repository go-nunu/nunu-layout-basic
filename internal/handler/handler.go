package handler

import (
	"github.com/go-nunu/nunu-layout-base/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

package handlers

import (
	"restApi/internal/app/logging"
)

type Handlers struct {
	logger logging.Logger
}

func NewHandlers(logger logging.Logger) *Handlers {
	return &Handlers{

		logger: logger,
	}
}

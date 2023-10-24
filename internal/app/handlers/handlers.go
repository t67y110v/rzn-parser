package handlers

import (
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	logger *logrus.Logger
}

func NewHandlers(logger *logrus.Logger) *Handlers {
	return &Handlers{

		logger: logger,
	}
}

package handlers

import (
	"restApi/internal/app/logging"
	"restApi/internal/app/store"
)

type Handlers struct {
	logger logging.Logger
	store  store.UserStore
}

func NewHandlers(store store.UserStore, logger logging.Logger) *Handlers {
	return &Handlers{
		store:  store,
		logger: logger,
	}
}

package handlers

import (
	"errors"
	"restApi/internal/app/logging"
	"restApi/internal/app/store"
)

var (
	errorIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errorThisUserIsNotAdmin       = errors.New("this user is not an admin")
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

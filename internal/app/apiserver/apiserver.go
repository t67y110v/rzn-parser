package apiserver

import (
	"restApi/internal/app/logging"
)

// запуск сервера
func Start(config *Config) error {

	logger := logging.GetLogger()
	server := newServer(config, logger)
	return server.router.Listen(config.BindAddr)
}

package apiserver

import (
	"restApi/internal/app/logging"
)

// запуск сервера
func Start(config *Config) error {

	logger, err := logging.NewLogger()
	if err != nil {
		return err
	}
	server := newServer(config, logger)
	return server.router.Listen(config.BindAddr)
}

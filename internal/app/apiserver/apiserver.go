package apiserver

import "github.com/sirupsen/logrus"

// запуск сервера
func Start(config *Config, logger *logrus.Logger) error {

	server := newServer(config, logger)
	return server.router.Listen(config.BindAddr)
}

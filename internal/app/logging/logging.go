package logging

import (

	//	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() (*logrus.Logger, error) {
	logger := logrus.New()

	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	return logger, nil
}

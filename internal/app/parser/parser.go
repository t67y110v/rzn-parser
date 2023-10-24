package parser

import (
	"restApi/internal/app/client"

	"github.com/sirupsen/logrus"
)

type Parser struct {
	Client client.Client
	Logger *logrus.Logger
}

func NewParser(l *logrus.Logger) *Parser {
	c := client.NewHTTPClient()
	return &Parser{
		Client: c,
		Logger: l,
	}
}

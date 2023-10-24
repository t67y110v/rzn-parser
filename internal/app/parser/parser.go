package parser

import "restApi/internal/app/client"

type Parser struct {
	Client client.Client
}

func NewParser() *Parser {
	c := client.NewHTTPClient()
	return &Parser{
		Client: c,
	}
}

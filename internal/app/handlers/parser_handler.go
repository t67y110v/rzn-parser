package handlers

import (
	"bytes"
	"encoding/json"

	parser "restApi/internal/app/parser"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleParser() fiber.Handler {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		count, err := parser.Parser(req.Login, req.Password)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		type resp struct {
			Result string `json:"result"`
		}
		res := &resp{}
		res.Result = count
		return c.JSON(res)
	}
}

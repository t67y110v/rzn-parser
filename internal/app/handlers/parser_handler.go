package handlers

import (
	"bytes"
	"encoding/json"

	"restApi/internal/app/handlers/requests"
	"restApi/internal/app/handlers/responses"
	parser "restApi/internal/app/parser"

	"github.com/gofiber/fiber/v2"
)

// @Summary Parser
// @Description pars site to get informaion about nr
// @Tags         Parser
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.ParserLogin  true "create new user"
// @Success 200 {object} responses.ParserResult
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /parser/parse [post]
func (h *Handlers) HandleParser() fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := &requests.ParserLogin{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		count, err := parser.Parser(req.Login, req.Password, req.Path)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		res := &responses.ParserResult{}
		res.Result = count
		return c.JSON(res)
	}
}

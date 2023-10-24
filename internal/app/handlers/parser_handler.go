package handlers

import (
	"net/http"

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
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		p := parser.NewParser(h.logger)
		count, err := p.Parse(req.Login, req.Password, req.Path, req.FileName, req.Monthly)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(&responses.ParserResult{
			Result: count,
		})
	}
}

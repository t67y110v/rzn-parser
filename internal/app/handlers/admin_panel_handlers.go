package handlers

import (
	"net/http"
	"restApi/internal/app/handlers/requests"
	"restApi/internal/app/handlers/responses"
	"restApi/internal/app/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary AdminAccess
// @Description checking that user have admin rights
// @Tags         Admin
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.EmailPassword  true "check adming rights"
// @Success 200 {object} responses.Result
// @Failure 400 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/access [post]
func (h *Handlers) HandleAdminAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &requests.EmailPassword{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			c.Status(http.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if !utils.CheckThatUserIsAdmin(u) {
			c.Status(http.StatusConflict)
			return c.JSON(fiber.Map{
				"message": "user is not an admin",
			})
		} else {

			res := &responses.Result{}

			res.Result = true
			return c.JSON(res)
		}

	}
}

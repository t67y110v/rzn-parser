package handlers

import (
	"bytes"
	"encoding/json"
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
// @Failure 500 {object} responses.Error
// @Router /admin/access [post]
func (h *Handlers) HandleAdminAccess() fiber.Handler {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		if !utils.CheckThatUserIsAdmin(u) {
			return c.JSON(fiber.Map{
				"message": err,
			})
		} else {

			res := &responses.Result{}

			res.Result = true
			return c.JSON(res)
		}

	}
}

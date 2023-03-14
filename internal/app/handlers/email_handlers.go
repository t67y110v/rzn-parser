package handlers

import (
	"bytes"
	"encoding/json"
	"restApi/internal/app/handlers/requests"
	"restApi/internal/app/handlers/responses"
	mail "restApi/internal/app/mailservice"

	"github.com/gofiber/fiber/v2"
)

// @Summary Email
// @Description send email
// @Tags         Email
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.EmailReq true "handler for sending message on email"
// @Success 200 {object} responses.Result
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /email/send [post]
func (h *Handlers) HandleSendEmail(emailSender, passwordSender, smtpEmail string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := &requests.EmailReq{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		err := mail.SendEmailMessage(emailSender, passwordSender, smtpEmail, req.RecipientMail, req.Subject, req.Body, h.logger)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		res := &responses.Result{}
		if err == nil {
			res.Result = true
		} else {
			res.Result = false
		}
		return c.JSON(res)
	}
}

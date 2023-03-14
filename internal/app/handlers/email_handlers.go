package handlers

import (
	"bytes"
	"encoding/json"
	mail "restApi/internal/app/mailservice"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleSendEmail(emailSender, passwordSender, smtpEmail string) fiber.Handler {
	type request struct {
		RecipientMail string `json:"recipient_mail"`
		Subject       string `json:"subject"`
		Body          string `json:"body"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
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
		type resp struct {
			Result bool `json:"result"`
		}
		res := &resp{}
		if err == nil {
			res.Result = true
		} else {
			res.Result = false
		}
		return c.JSON(res)
	}
}

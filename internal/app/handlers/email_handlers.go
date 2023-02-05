package handlers

import (
	"encoding/json"
	"net/http"
	mail "restApi/internal/app/mailservice"
)

func (h *Handlers) HandleSendEmail(emailSender, passwordSender, smtpEmail string) http.HandlerFunc {
	type request struct {
		RecipientMail string `json:"recipient_mail"`
		Subject       string `json:"subject"`
		Body          string `json:"body"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		err := mail.SendEmailMessage(emailSender, passwordSender, smtpEmail, req.RecipientMail, req.Subject, req.Body, h.logger)
		if err != nil {
			h.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			h.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
			return
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
		h.respond(w, r, http.StatusOK, res)
		h.logger.Infof("handle /sendEmail, status :%d", http.StatusOK)
	}
}

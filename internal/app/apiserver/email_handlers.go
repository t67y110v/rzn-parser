package apiserver

import (
	"encoding/json"
	"net/http"
	mail "restApi/internal/app/mailservice"
)

func (s *Server) handleSendEmail() http.HandlerFunc {
	type request struct {
		RecipientMail string `json:"recipient_mail"`
		Subject       string `json:"subject"`
		Body          string `json:"body"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		err := mail.SendEmailMessage(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail, req.RecipientMail, req.Subject, req.Body, s.logger)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
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
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /sendEmail, status :%d", http.StatusOK)
	}
}

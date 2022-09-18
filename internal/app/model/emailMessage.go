package model

type Email struct {
	RecipientMail string `json:"recipient_mail"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
}

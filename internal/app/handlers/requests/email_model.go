package requests

type EmailReq struct {
	RecipientMail string `json:"recipient_mail"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
}

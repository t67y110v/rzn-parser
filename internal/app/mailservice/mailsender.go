package mailservice

import (
	"fmt"
	"restApi/internal/app/logging"

	//"restApi/internal/app/model"

	gomail "gopkg.in/mail.v2"
)

func SendEmailMessage(sendersMail, sendersMailPassword, smtpEmail, RecipientMail, Subject, Body string, l logging.Logger) error {
	fmt.Println(sendersMail)
	fmt.Println(sendersMailPassword)
	message := gomail.NewMessage()
	message.SetHeader("From", sendersMail)
	message.SetHeader("To", RecipientMail)
	message.SetHeader("Subject", Subject)
	message.SetBody("text/html", Body) //
	a := gomail.NewDialer(smtpEmail, 465, sendersMail, sendersMailPassword)
	a.StartTLSPolicy = gomail.MandatoryStartTLS
	if err := a.DialAndSend(message); err != nil {
		l.Warningf("Dial and send MailSender error :%e", err)
		return err
	}
	return nil
}

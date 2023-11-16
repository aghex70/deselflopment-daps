package email

import (
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"time"
)

func SendMail(e domain.Email) error {
	from := mail.NewEmail(e.Source, e.From)
	subject := e.Subject
	to := mail.NewEmail(e.To, e.Recipient)
	message := mail.NewSingleEmail(
		from, subject, to, e.Body+time.Now().Format("2006-01-02 15:04:05"), e.Body)
	client := sendgrid.NewSendClient(pkg.SendGridApiKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return pkg.SendEmailError
	}
	return nil
}

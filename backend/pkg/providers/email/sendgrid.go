package email

import (
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/pkg"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(e domain.Email) error {
	from := mail.NewEmail(pkg.FromName, pkg.FromEmail)
	subject := e.Subject
	to := mail.NewEmail(e.Recipient, e.To)
	message := mail.NewSingleEmail(from, subject, to, "", e.Body)
	client := sendgrid.NewSendClient(pkg.SendGridApiKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}

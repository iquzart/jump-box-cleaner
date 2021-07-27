// send email using SendGrid
package notification

import (
	"bytes"
	"jump-box-cleaner/configs"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ViaSendGrid(cfg *configs.Config, body bytes.Buffer) {

	from := mail.NewEmail(cfg.NotifyEmail.FromName, cfg.NotifyEmail.FromEmail)
	subject := cfg.NotifyEmail.EmailSubject
	to := mail.NewEmail(cfg.NotifyEmail.ToName, cfg.NotifyEmail.ToEMail)
	htmlContent := body.String()
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(cfg.NotifyEmail.SendgridAPI)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)

	}
}

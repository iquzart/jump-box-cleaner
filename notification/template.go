// Notificaion functions
package notification

import (
	"bytes"
	"fmt"
	"html/template"
	"jump-box-cleaner/configs"
	"jump-box-cleaner/models"
	"log"
)

// construct email HTML email
func EmailNotification(cfg *configs.Config, results []models.Data) {

	t, err := template.ParseFiles("templates/email-template.html")
	if err != nil {
		log.Fatal(err)
	}

	var body bytes.Buffer
	if !cfg.NotifyEmail.Sendgrid {
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Disk Cleaner \n%s\n\n", mimeHeaders)))
	}

	t.Execute(&body, results)

	if cfg.NotifyEmail.Sendgrid {
		log.Println("Using SendGrid as email relay")
		ViaSendGrid(cfg, body)
	} else {
		log.Println("Please a relay services for email")
	}
}

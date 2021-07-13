// Notificaion functions
package notification

import (
	"bytes"
	"fmt"
	"jump-box-cleaner/configs"
	"jump-box-cleaner/models"
	"log"
	"text/template"
)

// construct email HTML email
func EmailNotification(cfg *configs.Config, results []models.ParentDirectories) {

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
		ViaSendGrid(cfg, body)
	}
}

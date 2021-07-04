// Notificaion functions
package notification

import (
	"jump-box-cleaner/models"
	"log"
	"os"
	"text/template"
)

// construct email HTML email
func EmailTemplate(results []models.ParentDirectories) {

	t, err := template.ParseFiles("templates/email-template.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(os.Stdout, results)

}

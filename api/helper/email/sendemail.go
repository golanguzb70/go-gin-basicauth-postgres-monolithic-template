package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
)

func SendEmail(to []string, subject string, cfg config.Config, htmlpath string, body any) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		log.Println(err)
		return err
	}

	var k bytes.Buffer
	err = t.Execute(&k, body)
	if err != nil {
		return err
	}

	if k.String() == "" {
		fmt.Println("Error buffer")
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("Subject: %s", subject) + mime + k.String())
	// Authentication.
	auth := smtp.PlainAuth("", cfg.SMTPEmail, cfg.SMTPEmailPass, cfg.SMTPHost)

	// Sending email.
	err = smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPEmail, to, msg)
	return err
}

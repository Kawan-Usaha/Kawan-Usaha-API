package lib

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"net/smtp"
)

type EmailData struct {
	Code      string
	FirstName string
	Subject   string
}

var Templates = map[string]string{
	"email_verification":   "verificationCode.html",
	"email_promotion":      "promotion.txt",
	"email_reset_password": "resetPassword.html",
}

func SendMailSingleReceiver(receiver string, data *EmailData, templateName string) {
	go func() {
		auth := smtp.PlainAuth(os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"))

		var body bytes.Buffer
		template, err := ParseTemplateDir("templates")
		if err != nil {
			log.Panic("Could not parse template", err)
		}

		template.ExecuteTemplate(&body, templateName, &data)

		msg := []byte(
			fmt.Sprintf(
				"From: %s\r\n"+
					"To: %s\r\n"+
					"Content-Type: text/html; charset=UTF-8\r\n"+
					"Subject: %s\r\n"+
					"\r\n"+
					"%s", os.Getenv("EMAIL_FROM"), receiver, data.Subject, body.String(),
			),
		)

		errsend := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_USER"), []string{receiver}, msg)

		if errsend != nil {
			log.Panic(err)
		}
	}()
}

func SendMailMultipleReceiver(receivers []string, data *EmailData, templateName string) {
	go func() {
		auth := smtp.PlainAuth(os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"))

		var body bytes.Buffer
		template, err := ParseTemplateDir("templates")
		if err != nil {
			log.Panic("Could not parse template", err)
		}

		template.ExecuteTemplate(&body, templateName, &data)

		msg := []byte(
			fmt.Sprintf(
				"From: %s\r\n"+
					"To: %s\r\n"+
					"Content-Type: text/html; charset=UTF-8\r\n"+
					"Subject: %s\r\n"+
					"\r\n"+
					"%s", os.Getenv("EMAIL_FROM"), strings.Join(receivers, ","), data.Subject, body.String(),
			),
		)

		errsend := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_USER"), receivers, msg)

		if errsend != nil {
			log.Panic(err)
		}
	}()
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

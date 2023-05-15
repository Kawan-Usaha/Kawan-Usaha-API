package lib

import (
	"log"
	"os"
	"strings"

	"net/smtp"
)

func SendMailMultipleReceiver(receiver []string, message []byte) {
	auth := smtp.PlainAuth(os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"))

	to := strings.Join(receiver, ",")

	msg := []byte("To: " + to + "\r\n" +
		"\r\n" +
		"Good afternoon!!!\r\n")
	err := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_USER"), receiver, msg)

	if err != nil {
		log.Fatal(err)
	}

}

func SendMailSingleReceiver(receiver string, message []byte) {
	auth := smtp.PlainAuth(os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"))

	msg := []byte("To: " + receiver + "\r\n" +
		"\r\n" +
		"Good afternoon!!!\r\n")
	err := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_USER"), []string{receiver}, msg)

	if err != nil {
		log.Fatal(err)
	}

}

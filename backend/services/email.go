package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(recipient string, subject string, body string) error {
	godotenv.Load()

	sender := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("SMTP_PROVIDER")
	port := os.Getenv("SMTP_PROVIDER_PORT")

	auth := smtp.PlainAuth(
		"",
		sender,
		password,
		host,
	)

	msg := "Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n"

	err := smtp.SendMail(
		host+":"+port,
		auth,
		sender,
		[]string{recipient},
		[]byte(msg),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

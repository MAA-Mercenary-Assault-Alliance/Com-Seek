package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
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
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func SendEmailWithAttachment(recipient string, subject string, body string, attachment *multipart.FileHeader) error {
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

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	textPart, err := writer.CreatePart(
		map[string][]string{"Content-Type": {"text/plain; charset=UTF-8"}},
	)
	if err != nil {
		return fmt.Errorf("failed to create text part: %w", err)
	}
	_, err = textPart.Write([]byte(body))
	if err != nil {
		return fmt.Errorf("failed to write body: %w", err)
	}

	file, err := attachment.Open()
	if err != nil {
		return fmt.Errorf("failed to open attachment file: %w", err)
	}
	defer file.Close()

	attachmentName := attachment.Filename
	attachmentType := getFileContentType(attachment)

	attachPart, err := writer.CreatePart(
		map[string][]string{
			"Content-Type":              {attachmentType},
			"Content-Disposition":       {fmt.Sprintf(`attachment; filename="%s"`, attachmentName)},
			"Content-Transfer-Encoding": {"base64"},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create attachment part: %w", err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, attachPart)
	_, err = io.Copy(encoder, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content to attachment: %w", err)
	}

	err = encoder.Close()
	if err != nil {
		return fmt.Errorf("failed to close base64 encoder: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	msg := fmt.Sprintf("Subject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n", subject, writer.Boundary())
	msg += buf.String()

	err = smtp.SendMail(
		host+":"+port,
		auth,
		sender,
		[]string{recipient},
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func getFileContentType(fileHeader *multipart.FileHeader) string {
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		return "application/octet-stream"
	}
	return contentType
}

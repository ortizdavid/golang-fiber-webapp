package helpers

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

func NewEmailService(smtpHost string, smtpPort int, username string, password string) *EmailService {
	return &EmailService{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
	}
}

func (es *EmailService) SendPlainEmail(to, subject, body string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	return es.sendEmail(to, message)
}


func (es *EmailService) SendHTMLEmail(to, subject, bodyHTML string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", to, subject, bodyHTML)
	return es.sendEmail(to, message)
}

func (es *EmailService) sendEmail(to, message string) error {
	auth := smtp.PlainAuth("", es.Username, es.Password, es.SMTPHost)
	addr := fmt.Sprintf("%s:%d", es.SMTPHost, es.SMTPPort)
	err := smtp.SendMail(addr, auth, es.Username, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	return nil
}
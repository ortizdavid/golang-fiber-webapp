package config

import (
	"os"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)

func DefaultEmailService() *helpers.EmailService {
	return &helpers.EmailService{
		SMTPHost: MailSMTPHost(),
		SMTPPort: MailSMTPPort(),
		Username: MailUser(),
		Password: MailPassword(),
	}
}

func MailUser() string {
	LoadDotEnv()
	return os.Getenv("MAIL_USER")
}

func MailPassword() string {
	LoadDotEnv()
	return os.Getenv("MAIL_PASSWORD")
}

func MailSMTPHost() string {
	LoadDotEnv()
	return os.Getenv("MAIL_SMTP_HOST")
}

func MailSMTPPort() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("MAIL_SMTP_PORT"))
}

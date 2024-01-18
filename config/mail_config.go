package config

import (
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
	return GetEnv("MAIL_USER")
}

func MailPassword() string {
	return GetEnv("MAIL_PASSWORD")
}

func MailSMTPHost() string {
	return GetEnv("MAIL_SMTP_HOST")
}

func MailSMTPPort() int {
	return helpers.ConvertToInt(GetEnv("MAIL_SMTP_PORT"))
}

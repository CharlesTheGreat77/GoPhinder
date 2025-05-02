package notifier

import (
	"fmt"
	"net/smtp"
)

type EmailNotifier struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
	To       string
}

// param constructor
func NewEmailNotifier(smtpHost string, smtpPort int, username string, password string, to string) *EmailNotifier {
	return &EmailNotifier{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
		To:       to,
	}
}

// send email notifications
func (e *EmailNotifier) Notify(title string, fields map[string]string) error {
	body := title + "\n\n"
	for k, v := range fields {
		body += fmt.Sprintf("%s: %s\n", k, v)
	}

	header := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n", e.To, title)
	msg := []byte(header + body)

	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)
	addr := fmt.Sprintf("%s:%d", e.SMTPHost, e.SMTPPort)
	return smtp.SendMail(addr, auth, e.Username, []string{e.To}, msg)
}

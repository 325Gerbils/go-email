package email

import (
	"net/smtp"
	"strings"

)

// Service contains data used to send emails
type Service struct {
	Email      string
	Pass       string
	Message    string
	Subject    string
	SMTPServer string
	SMTPPort   string
	Msg        string
	Auth       smtp.Auth
}

// Authenticate authenticates the email service
func Authenticate(user, pass string) *Service {
	e := &Service{}
	e.Email = user
	e.Pass = pass
	e.SMTPServer = "smtp.gmail.com"
	e.SMTPPort = ":587"
	auth := smtp.PlainAuth("", e.Email, e.Pass, e.SMTPServer)
	e.Auth = auth
	return e
}

// SetSubject sets the subject of the email
func (e *Service) SetSubject(subject string) {
	e.Subject = subject
	e.Msg += "Subject: " + subject + "\r\n"
}

// Write adds a line to the email body
func (e *Service) Write(line string) {
	e.Message += line + "\n"
}

// Send sends the email
func (e *Service) Send(to []string) error {

	e.Msg += "To: " + strings.Join(to, "\r\n")
	e.Msg += "Subject: " + e.Subject + "\r\n"
	e.Msg += e.Message

	msg := []byte(e.Msg)

	err := smtp.SendMail(e.SMTPServer+e.SMTPPort, e.Auth, e.Email, to, msg)
	if err != nil {
		return err
	}
	return nil
}

package email

import (
	"net/smtp"
	"strings"

)

var (
	authEmail  string
	authPass   string
	smtpServer string
	smtpPort   string
	plainAuth  smtp.Auth
)

// Auth initializes an smtp.PlainAuth variable
// Whether the authentication worked is not checked until an email is sent
func Auth(email, pass, smtpDomain string) {
	serverAndPort := strings.Split(smtpDomain, ":")
	smtpServer, smtpPort = serverAndPort[0], serverAndPort[1]
	plainAuth = smtp.PlainAuth("", email, pass, smtpServer)
}

// Send attempts to send an email with the specified content
// Will return an error if the email cannot be sent
// Todo: add attachments=[]string{filenames} as an argument
func Send(recip, subject, message string) error {
	var content string

	content += "To: " + recip + "\r\n"
	content += "Subject: " + subject + "\r\n"
	content += "Message: " + message + "\r\n"

	err := smtp.SendMail(smtpServer+":"+smtpPort, plainAuth, authEmail, []string{recip}, []byte(content))
	if err != nil {
		return err
	}

	return nil
}

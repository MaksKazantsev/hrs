package utils

import (
	"fmt"
	"net/smtp"
)

type CodeSender interface {
	SendCode(code, to string) error
}

func NewCodeSender() CodeSender {
	return &email{
		smtpHost:       "smtp.gmail.com",
		smtpPort:       587,
		senderEmail:    "emailfortestingservices@gmail.com",
		senderPassword: "thiqssqrlslemqsj",
	}
}

type email struct {
	smtpHost       string
	smtpPort       int
	senderEmail    string
	senderPassword string
}

func (e *email) SendCode(code, to string) error {
	msg := "Subject: Email verification message \n Your code: " + code
	auth := smtp.PlainAuth("", e.senderEmail, e.senderPassword, e.smtpHost)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", e.smtpHost, e.smtpPort), auth, e.senderEmail, []string{to}, []byte(msg)); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

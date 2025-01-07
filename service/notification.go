package service

import (
	"fmt"
	"os"

	"github.com/harlitad/notitication-service/pkg/email"
)

type TypeNotification string

const (
	TypeNotificationEmail TypeNotification = "email"
	TypeNotificationSMS   TypeNotification = "sms"
)

type ContactInfo struct {
	Username     string
	PhoneNumber  string
	EmailAddress string
}

// create interface factory
type NotificationFactory interface {
	Send(contact ContactInfo) string
}

// products
type SMSNotification struct{}

func (s *SMSNotification) Send(contact ContactInfo) string {
	return fmt.Sprintf("SMS sent to %s for %s", contact.PhoneNumber, contact.Username)
}

type EmailNotification struct{}

func (s *EmailNotification) Send(contact ContactInfo) string {
	// logic to send email
	emailService := email.SparkPostEmailService{
		APIKey: os.Getenv("MAILGUN_API_KEY"),
		Domain: os.Getenv("MAILGUN_DOMAIN"),
	}
	err := emailService.SendEmail(contact.EmailAddress)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("SMS sent to %s for %s", contact.EmailAddress, contact.Username)
}

func GetNotificationFactory(typeNotif string) (NotificationFactory, error) {
	if typeNotif == string(TypeNotificationSMS) {
		return &SMSNotification{}, nil
	} else if typeNotif == string(TypeNotificationEmail) {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("unknown notification type: %s", typeNotif)
}

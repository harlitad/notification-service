package service

import (
	"fmt"
	"os"

	"github.com/harlitad/notitication-service/pkg/email"
)

type ContactInfo struct {
	Username     string
	PhoneNumber  string
	EmailAddress string
}

// create interface
type Notification interface {
	Send(contact ContactInfo) string
}

// the products
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

// factory
type NotificationFactory interface {
	CreateNotification() Notification
}

// Concrete Factories
type SMSFactory struct{}

func (sf *SMSFactory) CreateNotification() Notification {
	return &SMSNotification{}
}

type EmailFactory struct{}

func (ef *EmailFactory) CreateNotification() Notification {
	return &EmailNotification{}
}

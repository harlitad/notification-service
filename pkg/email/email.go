package email

import (
	"log"

	"github.com/mailgun/mailgun-go"
)

type EmailSender interface {
	SendEmail() error
}

type SparkPostEmailService struct {
	APIKey string
	Domain string
}

func (s *SparkPostEmailService) SendEmail(to string) error {
	mg := mailgun.NewMailgun(s.Domain, s.APIKey)

	message := mg.NewMessage(
		"cc@"+s.Domain,
		"Hello",
		"Testing some Mailgun awesomeness!",
		to,
	)

	message.EnableTestMode()

	_, _, err := mg.Send(message)
	if err != nil {
		log.Printf("Failed to send email: %s\n", err)
		return err
	}

	log.Printf("Email sent successfully to %s\n", "harlitad")
	return nil
}

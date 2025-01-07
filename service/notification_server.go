package service

import (
	context "context"
	"fmt"

	grpc "google.golang.org/grpc"
)

type server struct {
	UnimplementedNotificationServiceServer
}

func (s *server) SendNotification(ctx context.Context, req *NotificationRequest) (*NotificationResponse, error) {
	contactInfo := &ContactInfo{
		Username:     req.GetUsername(),
		PhoneNumber:  req.GetPhoneNumber(),
		EmailAddress: req.GetEmailAddress(),
	}

	// mari menjahit
	var factory NotificationFactory
	if req.GetTypeNotification() == "sms" {
		factory = &SMSFactory{}
	} else if req.GetTypeNotification() == "email" {
		factory = &EmailFactory{}
	} else {
		return nil, fmt.Errorf("unknown notification type: %s", req.GetTypeNotification())
	}

	result := factory.CreateNotification().Send(*contactInfo)

	return &NotificationResponse{
		Message: result,
	}, nil

}

func NewNotificationServer(grpcServer *grpc.Server) {
	RegisterNotificationServiceServer(grpcServer, &server{})
}

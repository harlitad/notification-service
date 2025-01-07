package service

import (
	context "context"

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

	factory, _ := GetNotificationFactory(req.TypeNotification)

	result := factory.Send(*contactInfo)

	return &NotificationResponse{
		Message: result,
	}, nil

}

func NewNotificationServer(grpcServer *grpc.Server) {
	RegisterNotificationServiceServer(grpcServer, &server{})
}

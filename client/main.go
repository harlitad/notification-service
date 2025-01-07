package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/harlitad/notitication-service/service"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("server", "localhost:50051", "Address of the gRPC server")
	notificationType := flag.String("type", "", "Notification type (sms or email)")
	username := flag.String("username", "", "Username of the recipient")
	phoneNumber := flag.String("phone", "", "Phone number (required for SMS)")
	emailAddress := flag.String("email", "", "Email address (required for Email)")

	flag.Parse()

	if *notificationType != "sms" && *notificationType != "email" {
		log.Fatalf("Invalid or missing --type flag. Use 'sms' or 'email'.")
	}
	if *username == "" {
		log.Fatalf("--username is required.")
	}
	if *notificationType == "sms" && *phoneNumber == "" {
		log.Fatalf("--phone is required for SMS notifications.")
	}
	if *notificationType == "email" && *emailAddress == "" {
		log.Fatalf("--email is required for Email notifications.")
	}

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := service.NewNotificationServiceClient(conn)

	req := &service.NotificationRequest{
		TypeNotification: *notificationType,
		Username:         *username,
	}
	if *notificationType == "sms" {
		req.PhoneNumber = *phoneNumber
	} else {
		req.EmailAddress = *emailAddress
	}

	resp, err := client.SendNotification(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Println("Response from server:", resp.GetMessage())
}

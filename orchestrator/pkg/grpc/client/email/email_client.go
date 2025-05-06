package email

import (
	"context"
	"fmt"
	"log"
	"tiny-letter/orchestrator/pkg/app"
	pb_email_service "tiny-letter/orchestrator/pkg/grpc/pb/email_service"
	"tiny-letter/orchestrator/pkg/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	emailConn   *grpc.ClientConn
	emailClient pb_email_service.EmailServiceClient
)

type EmailClient struct{}

func initializeSubscriptionClient(addr string) error {
	var err error
	emailConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}
	emailClient = pb_email_service.NewEmailServiceClient(emailConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ConnectSubscriptionClient(grpcConfig *app.CommConfig) (*EmailClient, error) {
	subscriptionAddr := fmt.Sprintf("%s:%d", grpcConfig.Domain, grpcConfig.Port)
	emailConnErr := initializeSubscriptionClient(subscriptionAddr)
	if emailConnErr != nil {
		ShutdownSubscriptionClient()
		return nil, emailConnErr
	}
	return &EmailClient{}, nil
}

func ShutdownSubscriptionClient() {
	if emailConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		emailConn.Close()
	}
}
func OnboardUser(data *models.OnboardUserData) error {
	_, err := emailClient.OnboardUser(context.Background(), &pb_email_service.OnboardUserRequest{
		UserId: int32(data.UserId),
		Email:  data.Email,
		Role:   data.Role,
	})

	return err
}

package auth

import (
	"context"
	"fmt"
	"log"
	pb_auth_manager "tiny-letter/orchestrator/pkg/grpc/pb/auth_manager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authConn   *grpc.ClientConn
	authClient pb_auth_manager.AuthManagerClient
)

func InitializeAuthClient(addr string) error {
	var err error
	authConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	authClient = pb_auth_manager.NewAuthManagerClient(authConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ShutdownAuthClient() {
	if authConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		authConn.Close()
	}
}

func CreatePublisher(email, password string, planID int) (int, error) {
	res, err := authClient.CreatePublisher(context.Background(), &pb_auth_manager.CreatePublisherRequest{
		Email:    email,
		Password: password,
		PlanId:   int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(res.UserId), nil
}

func RollbackCreatePublisher(userId int) error {
	_, err := authClient.RollbackCreatePublisher(context.Background(), &pb_auth_manager.RollbackRequest{
		UserId: int32(userId),
	})
	if err != nil {
		return err
	}

	return nil
}

func CreateSubscriber(email, password string) (int, error) {
	res, err := authClient.CreateSubscriber(context.Background(), &pb_auth_manager.CreateSubscriberRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return -1, err
	}

	return int(res.UserId), nil
}

func RollbackCreateSubscriber(userId int) error {
	_, err := authClient.RollbackCreateSubscriber(context.Background(), &pb_auth_manager.RollbackRequest{
		UserId: int32(userId),
	})
	if err != nil {
		return err
	}

	return nil
}

package grpc_client

import (
	"context"
	"fmt"
	"log"
	pb_auth "tiny-letter/orchestrator/cmd/grpc/pb/auth"
	"tiny-letter/orchestrator/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authConn   *grpc.ClientConn
	authClient pb_auth.NotifyAuthClient
)

func initializeAuthClient() error {
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.Auth.Domain, config.Auth.Port)

	var err error
	authConn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	authClient = pb_auth.NewNotifyAuthClient(authConn)
	log.Println("Auth gRPC client successfully initialized.")
	return nil
}

func InitAuthClient() error {
	if authConn != nil {
		return initializeAuthClient()
	}
	return nil
}

func AuthPublisherAction(actionType pb_auth.PublisherActionType, userId, planId int) error {
	_, err := authClient.PublisherAction(context.Background(), &pb_auth.PublisherActionRequest{
		Action: actionType,
		UserId: int32(userId),
		PlanId: int32(planId),
	})
	if err != nil {
		log.Printf("Failed to execute PublisherAction: %v", err)
		return err
	}
	log.Printf("PublisherAction executed successfully for userId: %d, actionType: %v", userId, actionType)
	return nil
}

func AuthSubscriberAction(actionType pb_auth.SubscriberActionType, userId, planId, publicationId int) error {
	_, err := authClient.SubscriberAction(context.Background(), &pb_auth.SubscriberActionRequest{
		Action:        actionType,
		UserId:        int32(userId),
		PlanId:        int32(planId),
		PublicationId: int32(publicationId),
	})
	if err != nil {
		log.Printf("Failed to execute SubscriberAction: %v", err)
		return err
	}
	log.Printf("SubscriberAction executed successfully for userId: %d, actionType: %v", userId, actionType)
	return nil
}

func CheckAuthHealth() error {
	_, err := authClient.HealthCheck(context.Background(), &pb_auth.HealthCheckRequest{})
	return err
}

func ShutdownAuthClient() {
	if authConn != nil {
		log.Println("Closing Auth gRPC connection...")
		authConn.Close()
	}
}

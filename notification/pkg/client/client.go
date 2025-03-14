package client

import (
	"context"
	"fmt"
	"log"
	"tiny-letter/notification/pkg/app"
	pb_subscription_manager "tiny-letter/notification/pkg/grpc/pb/subscription_manager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	subscriptionConn   *grpc.ClientConn
	subscriptionClient pb_subscription_manager.SubscriptionManagerClient
)

type SubscriptionClient struct{}

func initializeSubscriptionClient(addr string) error {
	var err error
	subscriptionConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}
	subscriptionClient = pb_subscription_manager.NewSubscriptionManagerClient(subscriptionConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ConnectSubscriptionClient(grpcConfig *app.CommConfig) (*SubscriptionClient, error) {
	subscriptionAddr := fmt.Sprintf("%s:%d", grpcConfig.Domain, grpcConfig.Port)
	subscriptionConnErr := initializeSubscriptionClient(subscriptionAddr)
	if subscriptionConnErr != nil {
		ShutdownSubscriptionClient()
		return nil, subscriptionConnErr
	}
	return &SubscriptionClient{}, nil
}

func ShutdownSubscriptionClient() {
	if subscriptionConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		subscriptionConn.Close()
	}
}
func (s *SubscriptionClient) GetContentSubscribers(publicationId int) ([]int32, error) {
	resp, err := subscriptionClient.GetContentSubscribers(context.Background(), &pb_subscription_manager.GetContentSubscribersRequest{
		PublicationId: int32(publicationId),
	})
	if err != nil {
		return []int32{}, err
	}
	return resp.SubscriberIds, nil
}

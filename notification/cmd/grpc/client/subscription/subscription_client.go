package subscription

import (
	"context"
	"fmt"
	"log"
	pb_subscription_manager "tiny-letter/notification/cmd/grpc/pb/subscription_manager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	subscriptionConn   *grpc.ClientConn
	subscriptionClient pb_subscription_manager.SubscriptionManagerClient
)

func InitializeSubscriptionClient(addr string) error {
	var err error
	subscriptionConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	subscriptionClient = pb_subscription_manager.NewSubscriptionManagerClient(subscriptionConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ShutdownSubscriptionClient() {
	if subscriptionConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		subscriptionConn.Close()
	}
}

func GetContentSubscribers(publicationId int) ([]int32, error) {
	resp, err := subscriptionClient.GetContentSubscribers(context.Background(), &pb_subscription_manager.GetContentSubscribersRequest{
		PublicationId: int32(publicationId),
	})

	if err != nil {
		return []int32{}, err
	}

	return resp.SubscriberIds, nil
}

package subscription

import (
	"context"
	"fmt"
	"log"
	pb_subscription_manager "tiny-letter/orchestrator/pkg/grpc/pb/subscription_manager"

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

func CreateSubscriptionForPublisher(userID, planID int) (int, error) {
	resp, err := subscriptionClient.CreateSubscriptionForPublisher(context.Background(), &pb_subscription_manager.CreateSubscriptionForPublisherRequest{
		UserId: int32(userID),
		PlanId: int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(resp.SubscriptionId), nil
}

func RollbackCreateSubscriptionForPublisher(subscriptionID int) error {
	_, err := subscriptionClient.RollbackCreateSubscriptionForPublisher(context.Background(), &pb_subscription_manager.RollbackCreateSubscriptionForPublisherRequest{
		SubscriptionId: int32(subscriptionID),
	})
	if err != nil {
		return err
	}

	return nil
}

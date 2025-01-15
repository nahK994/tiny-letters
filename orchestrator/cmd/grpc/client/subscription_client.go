package grpc_client

import (
	"context"
	"fmt"
	"log"
	pb_subscription "tiny-letter/orchestrator/cmd/grpc/pb/subscription"
	"tiny-letter/orchestrator/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	subscriptionConn   *grpc.ClientConn
	subscriptionClient pb_subscription.NotifySubscriptionClient
)

func initializeSubscriptionClient() error {
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.Subscription.Domain, config.Subscription.Port)

	var err error
	subscriptionConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	subscriptionClient = pb_subscription.NewNotifySubscriptionClient(subscriptionConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func InitSubscriptionClient() error {
	if subscriptionConn == nil {
		return initializeSubscriptionClient()
	}
	return nil
}

func ShutdownSubscriptionClient() {
	if subscriptionConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		subscriptionConn.Close()
	}
}

func JoinPublication(userID, publicationID, subscriptionPlanID int) error {
	_, err := subscriptionClient.JoinPublication(context.Background(), &pb_subscription.JoinPublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
		PlanId:        int32(subscriptionPlanID),
	})
	if err != nil {
		return fmt.Errorf("JoinPublication failed: %v", err)
	}
	return nil
}

func LeavePublication(userID, publicationID int) error {
	_, err := subscriptionClient.LeavePublication(context.Background(), &pb_subscription.LeavePublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return fmt.Errorf("LeavePublication failed: %v", err)
	}
	return nil
}

func ChangePublicationPlan(userID, publicationID, newPlanID int) error {
	_, err := subscriptionClient.ChangePublicationPlan(context.Background(), &pb_subscription.ChangePublicationPlanRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
		NewPlanId:     int32(newPlanID),
	})
	if err != nil {
		return fmt.Errorf("ChangePublicationPlan failed: %v", err)
	}
	return nil
}

func SubscribePublisherPlan(userID, subscriptionPlanID int) error {
	_, err := subscriptionClient.SubscribePublisherPlan(context.Background(), &pb_subscription.SubscribePublisherRequest{
		UserId: int32(userID),
		PlanId: int32(subscriptionPlanID),
	})
	if err != nil {
		return fmt.Errorf("SubscribePublisherPlan failed: %v", err)
	}
	return nil
}

func UnsubscribePublisherPlan(userID int) error {
	_, err := subscriptionClient.UnsubscribePublisherPlan(context.Background(), &pb_subscription.UnsubscribePublisherRequest{
		UserId: int32(userID),
	})
	if err != nil {
		return fmt.Errorf("UnsubscribePublisherPlan failed: %v", err)
	}
	return nil
}

func ChangePublisherPlan(userID, newPlanID int) error {
	_, err := subscriptionClient.ChangePublisherPlan(context.Background(), &pb_subscription.ChangePublisherPlanRequest{
		UserId:    int32(userID),
		NewPlanId: int32(newPlanID),
	})
	if err != nil {
		return fmt.Errorf("ChangePublisherPlan failed: %v", err)
	}
	return nil
}

func BookPublisherSubscription(userID, planID int) error {
	_, err := subscriptionClient.BookPublisherSubscription(context.Background(), &pb_subscription.BookPublisherSubscriptionRequest{
		UserId: int32(userID),
		PlanId: int32(planID),
	})
	if err != nil {
		return fmt.Errorf("BookPublisherSubscription failed: %v", err)
	}
	return nil
}

func BookSubscriberSubscription(userID, publicationID, planID int) error {
	_, err := subscriptionClient.BookSubscriberSubscription(context.Background(), &pb_subscription.BookSubscriberSubscriptionRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
		PlanId:        int32(planID),
	})
	if err != nil {
		return fmt.Errorf("BookSubscriberSubscription failed: %v", err)
	}
	return nil
}

func HealthCheckSubscription() error {
	if _, err := subscriptionClient.HealthCheckSubscription(context.Background(), &pb_subscription.HealthCheckRequest{}); err != nil {
		return err
	}

	return nil
}

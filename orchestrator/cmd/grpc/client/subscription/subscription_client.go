package grpc_subscription

import (
	"context"
	"fmt"
	"log"
	pb_subscription "tiny-letter/orchestrator/cmd/grpc/pb/subscription"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	subscriptionConn   *grpc.ClientConn
	subscriptionClient pb_subscription.NotifySubscriptionClient
)

func InitializeSubscriptionClient(addr string) error {
	var err error
	subscriptionConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	subscriptionClient = pb_subscription.NewNotifySubscriptionClient(subscriptionConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ShutdownSubscriptionClient() {
	if subscriptionConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		subscriptionConn.Close()
	}
}

func ConfirmPublisherSubscription(userID, planID int) (int, error) {
	response, err := subscriptionClient.ConfirmPublisherSubscription(context.Background(), &pb_subscription.ConfirmPublisherSubscriptionRequest{
		UserId: int32(userID),
		PlanId: int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackConfirmPublisherSubscription(SubscriptionId int) error {
	_, err := subscriptionClient.RollbackConfirmPublisherSubscription(context.Background(), &pb_subscription.RollbackConfirmPublisherSubscriptionRequest{
		SubscriptionId: int32(SubscriptionId),
	})

	return err
}

func RevokePublisherSubscription(userId int) (int, error) {
	response, err := subscriptionClient.RevokePublisherSubscription(context.Background(), &pb_subscription.RevokePublisherSubscriptionRequest{
		UserId: int32(userId),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackRevokePublisherSubscription(subscriptionId int) error {
	_, err := subscriptionClient.RollbackRevokePublisherSubscription(context.Background(), &pb_subscription.RollbackRevokePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
	})

	return err
}

func ChangePublisherSubscription(userId, planID int) (int, error) {
	response, err := subscriptionClient.ChangePublisherSubscription(context.Background(), &pb_subscription.ChangePublisherSubscriptionRequest{
		UserId: int32(userId),
		PlanId: int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackChangePublisherSubscription(subscriptionId, oldPlanId int) error {
	_, err := subscriptionClient.RollbackChangePublisherSubscription(context.Background(), &pb_subscription.RollbackChangePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
		OldPlanId:      int32(oldPlanId),
	})

	return err
}

func JoinPublication(userID, publicationID int, isPremium bool) (int, error) {
	response, err := subscriptionClient.JoinPublication(context.Background(), &pb_subscription.JoinPublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
		IsPremium:     isPremium,
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackJoinPublication(subscriptionID int) error {
	_, err := subscriptionClient.RollbackJoinPublication(context.Background(), &pb_subscription.RollbackJoinPublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

func LeavePublication(userID, publicationID int) (int, error) {
	response, err := subscriptionClient.LeavePublication(context.Background(), &pb_subscription.LeavePublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackLeavePublication(subscriptionID int) error {
	_, err := subscriptionClient.RollbackLeavePublication(context.Background(), &pb_subscription.RollbackLeavePublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

func ChangePublicationPlan(userID, publicationID int) (int, error) {
	response, err := subscriptionClient.ChangePublicationPlan(context.Background(), &pb_subscription.ChangePublicationPlanRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackChangePublicationPlan(subscriptionID int) error {
	_, err := subscriptionClient.RollbackChangePublicationPlan(context.Background(), &pb_subscription.RollbackChangePublicationPlanRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

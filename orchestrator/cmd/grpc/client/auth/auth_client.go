package grpc_auth

import (
	"context"
	"fmt"
	"log"
	pb_auth "tiny-letter/orchestrator/cmd/grpc/pb/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authConn   *grpc.ClientConn
	authClient pb_auth.NotifyAuthClient
)

func InitializeAuthClient(addr string) error {
	var err error
	authConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	authClient = pb_auth.NewNotifyAuthClient(authConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func ShutdownAuthClient() {
	if authConn != nil {
		log.Println("Closing Subscription gRPC connection...")
		authConn.Close()
	}
}

func ConfirmPublisherSubscription(userID, planID, subscriptionId int) error {
	_, err := authClient.ConfirmPublisherSubscription(context.Background(), &pb_auth.ConfirmPublisherSubscriptionRequest{
		UserId:         int32(userID),
		PlanId:         int32(planID),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func RevokePublisherSubscription(userId, subscriptionId int) error {
	_, err := authClient.RevokePublisherSubscription(context.Background(), &pb_auth.RevokePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func ChangePublisherSubscription(userId, planID, subscriptionId int) error {
	_, err := authClient.ChangePublisherSubscription(context.Background(), &pb_auth.ChangePublisherSubscriptionRequest{
		PlanId:         int32(planID),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func JoinPublication(userID, publicationID, subscriptionId int, isPremium bool) error {
	_, err := authClient.JoinPublication(context.Background(), &pb_auth.JoinPublicationRequest{
		UserId:         int32(userID),
		PublicationId:  int32(publicationID),
		SubscriptionId: int32(subscriptionId),
		IsPremium:      isPremium,
	})
	if err != nil {
		return err
	}

	return nil
}

func LeavePublication(userID, publicationID, subscriptionId int) error {
	_, err := authClient.LeavePublication(context.Background(), &pb_auth.LeavePublicationRequest{
		UserId:         int32(userID),
		PublicationId:  int32(publicationID),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func ChangePublicationPlan(userID, publicationID, subscriptionId int) error {
	_, err := authClient.ChangePublicationPlan(context.Background(), &pb_auth.ChangePublicationPlanRequest{
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

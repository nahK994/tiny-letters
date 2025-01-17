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

func RollbackConfirmPublisherSubscription(SubscriptionId int) error {
	_, err := authClient.RollbackConfirmPublisherSubscription(context.Background(), &pb_auth.RollbackConfirmPublisherSubscriptionRequest{
		SubscriptionId: int32(SubscriptionId),
	})

	return err
}

func RevokePublisherSubscription(userId, subscriptionId int) error {
	_, err := authClient.RevokePublisherSubscription(context.Background(), &pb_auth.RevokePublisherSubscriptionRequest{
		UserId:         int32(userId),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func RollbackRevokePublisherSubscription(subscriptionId int) error {
	_, err := authClient.RollbackRevokePublisherSubscription(context.Background(), &pb_auth.RollbackRevokePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
	})

	return err
}

func ChangePublisherSubscription(userId, planID, subscriptionId int) error {
	_, err := authClient.ChangePublisherSubscription(context.Background(), &pb_auth.ChangePublisherSubscriptionRequest{
		UserId:         int32(userId),
		PlanId:         int32(planID),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func RollbackChangePublisherSubscription(subscriptionId, oldPlanId int) error {
	_, err := authClient.RollbackChangePublisherSubscription(context.Background(), &pb_auth.RollbackChangePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
		OldPlanId:      int32(oldPlanId),
	})

	return err
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

func RollbackJoinPublication(subscriptionID int) error {
	_, err := authClient.RollbackJoinPublication(context.Background(), &pb_auth.RollbackJoinPublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
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

func RollbackLeavePublication(subscriptionID int) error {
	_, err := authClient.RollbackLeavePublication(context.Background(), &pb_auth.RollbackLeavePublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

func ChangePublicationPlan(userID, publicationID, subscriptionId int) error {
	_, err := authClient.ChangePublicationPlan(context.Background(), &pb_auth.ChangePublicationPlanRequest{
		UserId:         int32(userID),
		PublicationId:  int32(publicationID),
		SubscriptionId: int32(subscriptionId),
	})
	if err != nil {
		return err
	}

	return nil
}

func RollbackChangePublicationPlan(subscriptionID int) error {
	_, err := authClient.RollbackChangePublicationPlan(context.Background(), &pb_auth.RollbackChangePublicationPlanRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

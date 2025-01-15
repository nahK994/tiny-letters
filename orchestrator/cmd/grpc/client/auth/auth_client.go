package grpc_auth

import (
	"context"
	"fmt"
	"log"
	"time"
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

func CheckAvailability(client pb_auth.NotifyAuthClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb_auth.CheckAvailabilityRequest{}

	response, err := client.CheckAvailability(ctx, request)
	if err != nil {
		return err
	}

	log.Printf("CheckAvailability response: Success=%v", response.IsSuccess)
	return nil
}

func ConfirmPublisherSubscription(userID, planID int) (int, error) {
	response, err := authClient.ConfirmPublisherSubscription(context.Background(), &pb_auth.ConfirmPublisherSubscriptionRequest{
		UserId: int32(userID),
		PlanId: int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackConfirmPublisherSubscription(SubscriptionId int) error {
	_, err := authClient.RollbackConfirmPublisherSubscription(context.Background(), &pb_auth.RollbackConfirmPublisherSubscriptionRequest{
		SubscriptionId: int32(SubscriptionId),
	})

	return err
}

func RevokePublisherSubscription(userId int) (int, error) {
	response, err := authClient.RevokePublisherSubscription(context.Background(), &pb_auth.RevokePublisherSubscriptionRequest{
		UserId: int32(userId),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackRevokePublisherSubscription(subscriptionId int) error {
	_, err := authClient.RollbackRevokePublisherSubscription(context.Background(), &pb_auth.RollbackRevokePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
	})

	return err
}

func ChangePublisherSubscription(userId, planID int) (int, error) {
	response, err := authClient.ChangePublisherSubscription(context.Background(), &pb_auth.ChangePublisherSubscriptionRequest{
		UserId: int32(userId),
		PlanId: int32(planID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackChangePublisherSubscription(subscriptionId, oldPlanId int) error {
	_, err := authClient.RollbackChangePublisherSubscription(context.Background(), &pb_auth.RollbackChangePublisherSubscriptionRequest{
		SubscriptionId: int32(subscriptionId),
		OldPlanId:      int32(oldPlanId),
	})

	return err
}

func JoinPublication(userID, publicationID int, isPremium bool) (int, error) {
	response, err := authClient.JoinPublication(context.Background(), &pb_auth.JoinPublicationRequest{
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
	_, err := authClient.RollbackJoinPublication(context.Background(), &pb_auth.RollbackJoinPublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

func LeavePublication(userID, publicationID int) (int, error) {
	response, err := authClient.LeavePublication(context.Background(), &pb_auth.LeavePublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackLeavePublication(subscriptionID int) error {
	_, err := authClient.RollbackLeavePublication(context.Background(), &pb_auth.RollbackLeavePublicationRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

func ChangePublicationPlan(userID, publicationID int) (int, error) {
	response, err := authClient.ChangePublicationPlan(context.Background(), &pb_auth.ChangePublicationPlanRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return -1, err
	}

	return int(response.SubscriptionId), nil
}

func RollbackChangePublicationPlan(subscriptionID int) error {
	_, err := authClient.RollbackChangePublicationPlan(context.Background(), &pb_auth.RollbackChangePublicationPlanRequest{
		SubscriptionId: int32(subscriptionID),
	})

	return err
}

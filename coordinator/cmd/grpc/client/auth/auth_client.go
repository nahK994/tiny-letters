package grpc_auth

import (
	"context"
	"fmt"
	"log"
	pb_auth "tiny-letter/coordinator/cmd/grpc/pb/auth"

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

func ConfirmPublisherSubscription(userID, planID int) error {
	_, err := authClient.ConfirmPublisherSubscription(context.Background(), &pb_auth.ConfirmPublisherSubscriptionRequest{
		UserId: int32(userID),
		PlanId: int32(planID),
	})
	if err != nil {
		return err
	}

	return nil
}

func RevokePublisherSubscription(userId, planId int) error {
	_, err := authClient.RevokePublisherSubscription(context.Background(), &pb_auth.RevokePublisherSubscriptionRequest{
		UserId: int32(userId),
		PlanId: int32(planId),
	})
	if err != nil {
		return err
	}

	return nil
}

func ChangePublisherSubscription(userId, planID int) error {
	_, err := authClient.ChangePublisherSubscription(context.Background(), &pb_auth.ChangePublisherSubscriptionRequest{
		PlanId: int32(planID),
		UserId: int32(userId),
	})
	if err != nil {
		return err
	}

	return nil
}

func JoinPublication(userID, publicationID int, isPremium bool) error {
	_, err := authClient.JoinPublication(context.Background(), &pb_auth.JoinPublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
		IsPremium:     isPremium,
	})
	if err != nil {
		return err
	}

	return nil
}

func LeavePublication(userID, publicationID int) error {
	_, err := authClient.LeavePublication(context.Background(), &pb_auth.LeavePublicationRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return err
	}

	return nil
}

func ChangeSubscriberSubscription(userID, publicationID int) error {
	_, err := authClient.ChangeSubscriberSubscription(context.Background(), &pb_auth.ChangeSubscriberSubscriptionRequest{
		UserId:        int32(userID),
		PublicationId: int32(publicationID),
	})
	if err != nil {
		return err
	}

	return nil
}

package subscription

import (
	"context"
	"fmt"
	"log"
	pb_subscription "tiny-letter/content/cmd/grpc/pb/subscription"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	subscriptionConn   *grpc.ClientConn
	subscriptionClient pb_subscription.AskSubscriptionClient
)

func InitializeSubscriptionClient(addr string) error {
	var err error
	subscriptionConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Subscription service: %v", err)
	}

	subscriptionClient = pb_subscription.NewAskSubscriptionClient(subscriptionConn)
	log.Println("Subscription gRPC client successfully initialized.")
	return nil
}

func CheckIsAuthorizedPublisher(userID, publicationID int) (bool, error) {
	response, err := subscriptionClient.IsAuthorizedPublisher(context.Background(), &pb_subscription.IsAuthorizedPublisherRequest{
		PublicationId: int32(publicationID),
		UserId:        int32(userID),
	})
	if err != nil {
		return false, err
	}

	return response.IsSuccess, nil
}

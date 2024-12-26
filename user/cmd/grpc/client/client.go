package grpc_client

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "tiny-letter-user/cmd/grpc/pb/user-subscriber"
	"tiny-letter-user/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NotifySubscription(id, subscriptionType int, role string) error {
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%s", config.App.GRPC.User_Subscriber.Domain, config.App.GRPC.User_Subscriber.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Create(ctx, &pb.Request{Id: int32(id), Role: role, SubscriptionType: int32(subscriptionType)})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}

	if err != nil || !resp.IsSuccess {
		return fmt.Errorf("cannot create row in subscription table")
	}
	return nil
}

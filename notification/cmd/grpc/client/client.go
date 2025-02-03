package grpc_client

import (
	"fmt"
	"tiny-letter/notification/cmd/grpc/client/subscription"
	"tiny-letter/notification/pkg/app"
)

func IsGRPC_ClientAvailable(grpcConfig *app.CommConfig) error {
	subscriptionAddr := fmt.Sprintf("%s:%d", grpcConfig.Domain, grpcConfig.Port)
	subscriptionConnErr := subscription.InitializeSubscriptionClient(subscriptionAddr)
	if subscriptionConnErr != nil {
		subscription.ShutdownSubscriptionClient()
		return subscriptionConnErr
	}

	return nil
}

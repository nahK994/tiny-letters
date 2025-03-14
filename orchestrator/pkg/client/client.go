package client

import (
	"fmt"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/client/auth"
	"tiny-letter/orchestrator/pkg/client/subscription"
)

func ConnectGRPC(grpcConfig *app.GRPC) error {
	authAddr := fmt.Sprintf("%s:%d", grpcConfig.Auth.Domain, grpcConfig.Auth.Port)
	authConnErr := auth.InitializeAuthClient(authAddr)
	if authConnErr != nil {
		return authConnErr
	}

	subscriptionAddr := fmt.Sprintf("%s:%d", grpcConfig.Subscription.Domain, grpcConfig.Subscription.Port)
	subscriptionConnErr := subscription.InitializeSubscriptionClient(subscriptionAddr)
	if subscriptionConnErr != nil {
		auth.ShutdownAuthClient()
		return subscriptionConnErr
	}

	return nil
}

package grpc_client

import (
	"fmt"
	"tiny-letter/coordinator/cmd/grpc/client/auth"
	"tiny-letter/coordinator/cmd/grpc/client/subscription"
	"tiny-letter/coordinator/pkg/app"
)

func IsGRPC_ClientAvailable(grpcConfig *app.GRPC) error {
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

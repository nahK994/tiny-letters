package grpc_client

import (
	"fmt"
	"tiny-letter/content/cmd/grpc/client/subscription"
	"tiny-letter/content/pkg/app"
)

func IsGRPC_ClientAvailable(grpcConfig *app.GRPC_config) error {
	subsAddr := fmt.Sprintf("%s:%d", grpcConfig.Subscription.Domain, grpcConfig.Subscription.Port)
	if subsConnErr := subscription.InitializeSubscriptionClient(subsAddr); subsConnErr != nil {
		return subsConnErr
	}

	return nil
}

package grpc_client

import (
	grpc_auth "tiny-letter/coordinator/cmd/grpc/client/auth"
	grpc_subscription "tiny-letter/coordinator/cmd/grpc/client/subscription"
)

func IsGRPC_ClientAvailable(addr string) error {
	if err := grpc_auth.InitializeAuthClient(addr); err != nil {
		grpc_auth.ShutdownAuthClient()
		return err
	}

	if err := grpc_subscription.InitializeSubscriptionClient(addr); err != nil {
		grpc_subscription.ShutdownSubscriptionClient()
		return err
	}
	return nil
}

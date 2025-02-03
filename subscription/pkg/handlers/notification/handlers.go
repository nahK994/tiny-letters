package notification_handlers

import (
	"context"
	pb_notification "tiny-letter/subscription/cmd/grpc/pb/notification"
	"tiny-letter/subscription/pkg/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationListener struct {
	pb_notification.UnimplementedNotificationListenerServer
	db *db.Repository
}

func GetNotificationHandlers(db *db.Repository) *NotificationListener {
	return &NotificationListener{
		db: db,
	}
}

func (n *NotificationListener) GetContentSubscribers(c context.Context, req *pb_notification.GetContentSubscribersRequest) (*pb_notification.GetContentSubscribersResponse, error) {
	data := &db.GetContentSubscribersRequest{
		PublicationId: int(req.PublicationId),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	subscriberIds, err := n.db.GetContentSubscribers(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get subscriberIds: %v", err)
	}

	return &pb_notification.GetContentSubscribersResponse{
		SubscriberIds: subscriberIds,
	}, nil
}

package content_handlers

import (
	"context"
	pb_content "tiny-letter/subscription/cmd/grpc/pb/content"
	"tiny-letter/subscription/pkg/db"
)

type ContentListener struct {
	pb_content.UnimplementedContentListenerServer
	db *db.Repository
}

func GetContentHandlers(db *db.Repository) *ContentListener {
	return &ContentListener{
		db: db,
	}
}

func (l *ContentListener) IsAuthorizedPublisher(c context.Context, req *pb_content.IsAuthorizedPublisherRequest) (*pb_content.Response, error) {
	data := &db.IsAuthorizedPublisherRequest{
		UserId:        int(req.UserId),
		PublicationId: int(req.PublicationId),
	}
	if err := data.Validate(); err != nil {
		return nil, err
	}

	authorized, err := l.db.IsAuthorizedPublisher(data)
	if err != nil {
		return nil, err
	}

	return &pb_content.Response{IsAuthorized: authorized}, nil
}

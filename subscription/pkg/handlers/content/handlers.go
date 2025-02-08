package content_handlers

import (
	"context"
	pb_publication_authorization "tiny-letter/subscription/cmd/grpc/pb/publication_authorization"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/models"
)

type ContentListener struct {
	pb_publication_authorization.UnimplementedPublicationAuthorizationServer
	db *db.Repository
}

func GetContentHandlers(db *db.Repository) *ContentListener {
	return &ContentListener{
		db: db,
	}
}

func (l *ContentListener) IsAuthorizedPublisher(c context.Context, req *pb_publication_authorization.IsAuthorizedPublisherRequest) (*pb_publication_authorization.Response, error) {
	data := &models.IsAuthorizedPublisherRequest{
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

	return &pb_publication_authorization.Response{IsAuthorized: authorized}, nil
}

package publication_authorization_handlers

import (
	"context"
	pb_publication_authorization "tiny-letter/subscription/cmd/grpc/pb/publication_authorization"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/models"
)

type PublicationAuthorizationHandler struct {
	pb_publication_authorization.UnimplementedPublicationAuthorizationServer
	db *db.Repository
}

func GetPublicationAuthorizationHandler(db *db.Repository) *PublicationAuthorizationHandler {
	return &PublicationAuthorizationHandler{
		db: db,
	}
}

func (l *PublicationAuthorizationHandler) IsAuthorizedPublisher(c context.Context, req *pb_publication_authorization.IsAuthorizedPublisherRequest) (*pb_publication_authorization.Response, error) {
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

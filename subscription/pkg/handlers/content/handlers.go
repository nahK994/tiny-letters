package content_handlers

import (
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

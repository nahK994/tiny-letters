package content_handlers

import pb_content "tiny-letter/subscription/cmd/grpc/pb/content"

type ContentListener struct {
	pb_content.UnimplementedContentListenerServer
}

func GetEmailHandlers() *ContentListener {
	return &ContentListener{}
}

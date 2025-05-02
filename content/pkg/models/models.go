package models

import (
	"encoding/json"
	"fmt"
)

type Publication struct {
	Name        string
	PublisherID int
}

type Post struct {
	Title         string
	Content       string
	PublicationID int
	IsPremium     bool
}

type PublishContentData struct {
	ContentId     int     `json:"content_id"`
	Content       string  `json:"content"`
	Title         string  `json:"title"`
	SubscriberIds []int32 `json:"subscriber_ids"`
}

type MessageItem struct {
	Topic string          `json:"topic"`
	Data  json.RawMessage `json:"data"`
}

func (req *Publication) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if req.PublisherID <= 0 {
		return fmt.Errorf("publisher ID cannot be or negative")
	}

	return nil
}

func (req *Post) Validate() error {
	if req.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if req.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	if req.PublicationID <= 0 {
		return fmt.Errorf("publication ID cannot be or negative")
	}

	return nil
}

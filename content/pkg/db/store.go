package db

import (
	"fmt"
	"tiny-letter/content/pkg/models"
)

type (
	PublicationID int
	PostID        int
)

func (r Repository) CreatePublication(name string, publisherID int) (PublicationID, error) {
	var id PublicationID
	data := r.DB.QueryRow("INSERT INTO publications (name, publisher_id) VALUES ($1, $2) RETURNING id", name, publisherID)
	err := data.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("could not create publication: %w", err)
	}

	return id, nil
}

func (r Repository) CreatePost(title, content string, publicationID int, isPremium bool) (PostID, error) {
	var id PostID
	data := r.DB.QueryRow("INSERT INTO posts (title, content, publication_id, is_premium) VALUES ($1, $2, $3, $4) RETURNING id", title, content, publicationID, isPremium)
	err := data.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("could not create post: %w", err)
	}

	return id, nil
}

func (r *Repository) GetContentInfo(postID int) (*models.Post, error) {
	var post models.Post
	err := r.DB.QueryRow("SELECT title, content, is_premium, publication_id FROM posts WHERE id = $1", postID).Scan(&post.Title, &post.Content, &post.IsPremium, &post.PublicationID)
	if err != nil {
		return nil, fmt.Errorf("could not get content info: %w", err)
	}
	return &post, nil
}

func (r *Repository) MarkPostAsPublished(postID int) error {
	_, err := r.DB.Exec("UPDATE posts SET is_published = TRUE WHERE id = $1", postID)
	if err != nil {
		return fmt.Errorf("could not mark post as published: %w", err)
	}
	return nil
}

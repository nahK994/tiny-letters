package db

import "fmt"

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

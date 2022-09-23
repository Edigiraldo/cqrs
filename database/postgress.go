package database

import (
	"context"
	"database/sql"

	"github.com/Edigiraldo/cqrs/models"
	_ "github.com/lib/pq"
)

type PostgressRepository struct {
	db *sql.DB
}

func NewPostgressRepository(url string) (*PostgressRepository, error) {
	db, err := sql.Open("postgress", url)
	if err != nil {
		return nil, err
	}

	return &PostgressRepository{db}, err
}

func (repo *PostgressRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgressRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	query := "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)"
	_, err := repo.db.ExecContext(ctx, query, feed.Id, feed.Title, feed.Description, feed.CreatedAt)

	return err
}

func (repo *PostgressRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	query := "SELECT id, title, description, created_at FROM feeds"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := []*models.Feed{}
	for rows.Next() {
		feed := models.Feed{}
		err := rows.Scan(&feed.Id, &feed.Title, &feed.Description, &feed.CreatedAt)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, &feed)
	}

	return feeds, nil
}

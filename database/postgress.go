package database

import (
	"context"
	"database/sql"

	"github.com/Edigiraldo/cqrs/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, err
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	query := "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)"
	_, err := repo.db.ExecContext(ctx, query, feed.Id, feed.Title, feed.Description, feed.CreatedAt)

	return err
}

func (repo *PostgresRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
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

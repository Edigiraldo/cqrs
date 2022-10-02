package events

import (
	"context"

	"github.com/Edigiraldo/cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SuscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	OnCreateFeed(f func(CreatedFeedMessage)) error
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func Close() {
	eventStore.Close()
}

func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)
}

func SuscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	return eventStore.SuscribeCreatedFeed(ctx)
}

func OnCreateFeed(f func(CreatedFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}

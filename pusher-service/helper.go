package main

import "time"

type CreatedFeedMessage struct {
	Type        string    `json:"type"`
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func newCreatedFeedMessage(id, title, description string, createdAt time.Time) *CreatedFeedMessage {
	return &CreatedFeedMessage{
		Type:        "created_feed",
		Id:          id,
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}
}

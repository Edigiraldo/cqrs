package search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/Edigiraldo/cqrs/models"
	elastic "github.com/elastic/go-elasticsearch/v7"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElasticRepository(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, err
	}

	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {

}

func (r *ElasticRepository) IndexFeed(ctx context.Context, feed models.Feed) error {
	body, _ := json.Marshal(feed)
	_, err := r.client.Index(
		"feeds",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(feed.Id),
		r.client.Index.WithContext(ctx),
		r.client.Index.WithRefresh("wait_for"),
	)

	return err
}

func (r *ElasticRepository) SearchFeed(ctx context.Context, query string) (results []models.Feed, err error) {
	var buf bytes.Buffer

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"title", "description"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, err
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("feeds"),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			results = nil
		}
	}()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var eRes map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&eRes); err != nil {
		return nil, err
	}

	var feeds []models.Feed
	for _, hit := range eRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
		feed := models.Feed{}
		source := hit.(map[string]interface{})["_source"]
		marshal, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(marshal, &feed); err == nil {
			feeds = append(feeds, feed)
		}
	}

	return feeds, nil
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Edigiraldo/cqrs/database"
	"github.com/Edigiraldo/cqrs/events"
	"github.com/Edigiraldo/cqrs/repository"
	"github.com/Edigiraldo/cqrs/search"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPasword      string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticSearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

func main() {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	addr := fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		conf.PostgresUser,
		conf.PostgresPasword,
		conf.PostgresDB,
	)

	repo, err := database.NewPostgresRepository(addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	repository.SetRepository(repo)

	es, err := search.NewElasticRepository(fmt.Sprintf("http://%s", conf.ElasticSearchAddress))
	if err != nil {
		log.Fatalf("%v", err)
	}
	search.SetSearchRepository(es)
	defer search.Close()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", conf.NatsAddress))
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = n.OnCreateFeed(onCreateFeed)
	if err != nil {
		log.Fatalf("%v", err)
	}

	events.SetEventStore(n)
	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("%v", err)
	}
}

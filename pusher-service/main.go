package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Edigiraldo/cqrs/events"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	hub := NewHub()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", conf.NatsAddress))
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	err = n.OnCreateFeed(func(m events.CreatedFeedMessage) {
		hub.Broadcast(newCreatedFeedMessage(m.Id, m.Title, m.Description, m.CreatedAt), nil)
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	events.SetEventStore(n)
	defer events.Close()

	go hub.Run()

	http.HandleFunc("/ws", hub.HandleWebSocket)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

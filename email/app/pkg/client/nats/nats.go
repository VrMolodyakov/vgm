package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NewJetStream(url string, maxAsyncPending int) nats.JetStreamContext {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatalf("Failed while creating Stream: %v\n", err)
	}
	return js

}

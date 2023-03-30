package nats

import "github.com/nats-io/nats.go"

type publisher struct {
	stream nats.JetStreamContext
}

// NewPublisher Nats publisher constructor
func NewPublisher(stream nats.JetStreamContext) *publisher {
	return &publisher{stream: stream}
}

// Publish Publish will publish to the cluster and wait for an ACK
func (p *publisher) Publish(subject string, data []byte) error {
	_, err := p.stream.Publish(subject, data)
	return err
}

// PublishAsync PublishAsync will publish to the cluster and asynchronously process the ACK or error state.
func (p *publisher) PublishAsync(subject string, data []byte) (nats.PubAckFuture, error) {
	return p.stream.PublishAsync(subject, data)
}

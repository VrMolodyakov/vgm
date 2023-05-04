package nats

import (
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/nats-io/nats.go"
)

const (
	maxDelay = time.Second * 5
	attempts = 5
)

func NewStreamContext(
	host string,
	port int,
	subjectName string,
	subjects []string,

) (*nats.Conn, nats.JetStreamContext, error) {

	address := fmt.Sprintf("nats://%s:%d", host, port)
	var streamContext nats.JetStreamContext
	var connection *nats.Conn
	if err := retry.Do(func() error {
		fmt.Println("start attempt to get connection")
		connection, err := nats.Connect(address)
		if err != nil {
			fmt.Println(err)
			return err
		}
		streamContext, err = connection.JetStream(nats.PublishAsyncMaxPending(256))
		if err != nil {
			return err
		}
		_, err = streamContext.AddStream(&nats.StreamConfig{
			Name:     subjectName,
			Subjects: subjects,
		})
		if err != nil {
			return err
		}
		return nil

	}, retry.Delay(maxDelay), retry.Attempts(attempts)); err != nil {
		return nil, nil, err
	}
	return connection, streamContext, nil
}

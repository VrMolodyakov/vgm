package nats

import "time"

const (
	durableName = "emailservice-dur"
	ackWait     = 60 * time.Second
	maxInflight = 30
	maxDeliver  = 3
)

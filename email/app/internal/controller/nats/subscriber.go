package nats

import (
	"context"
	"sync"

	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"github.com/nats-io/nats.go"
)

type MsgHandler func(m *nats.Msg)

type EmailService interface {
	Send(email *model.Email) error
}

type subscriber struct {
	stream       nats.JetStreamContext
	emailService EmailService
	logger       logging.Logger
}

func NewSubscriber(
	stream nats.JetStreamContext,
	emailService EmailService,
	logger logging.Logger) *subscriber {

	return &subscriber{
		stream:       stream,
		emailService: emailService,
		logger:       logger,
	}

}

func (s *subscriber) Subscribe(subject, qgroup string, workersNum int, handler nats.MsgHandler) {
	s.logger.Infof("Subscribing to Subject: %v, group: %v", subject, qgroup)
	wg := &sync.WaitGroup{}
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		s.runWorker(wg, i, subject, qgroup, handler, nats.Durable(durableName))
	}
	wg.Wait()
}

func (s *subscriber) Run(ctx context.Context) {

}

func (s *subscriber) runWorker(
	wg *sync.WaitGroup,
	workerID int,
	subject string,
	qgroup string,
	handler nats.MsgHandler,
	opts ...nats.SubOpt) {

	s.logger.Infof("Subscribing worker: %v, subject: %v, qgroup: %v", workerID, subject, qgroup)
	defer wg.Done()
	sub, err := s.stream.QueueSubscribe(subject, qgroup, handler, opts...)
	if err != nil {
		s.logger.Errorf("WorkerID: %v, QueueSubscribe: %v", workerID, err)
		if err := sub.Unsubscribe(); err != nil {
			s.logger.Errorf("WorkerID: %v, conn.Close error: %v", workerID, err)
		}
	}
}

func (s *subscriber) processSendEmail(ctx context.Context) nats.MsgHandler {
	// s.stream.Ac
	return func(m *nats.Msg) {

	}
}

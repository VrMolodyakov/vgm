package nats

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/model"
	"github.com/VrMolodyakov/vgm/email/app/pkg/errors"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"github.com/avast/retry-go"
	"github.com/nats-io/nats.go"
)

const (
	retryAttempts = 3
	retryDelay    = 1 * time.Second
)

type MsgHandler func(m *nats.Msg)

type EmailUseCase interface {
	Send(ctx context.Context, email *model.Email) error
}

type SubscriberCfg struct {
	DurableName             string
	DeadMessageQueueSubject string
	SendEmailSubject        string
	EmailGroupName          string
	AckWait                 time.Duration
	MaxInflight             int
	SendWorkers             int
	MaxDeliver              int
}

func NewSubscriberCfg(
	durableName string,
	deadMessageQueueSubject string,
	sendEmailSubject string,
	emailGroupName string,
	ackWait time.Duration,
	sendWorkers int,
	maxInflight int,
	maxDeliver int,
) SubscriberCfg {

	return SubscriberCfg{
		DurableName:             durableName,
		DeadMessageQueueSubject: deadMessageQueueSubject,
		SendEmailSubject:        sendEmailSubject,
		EmailGroupName:          emailGroupName,
		AckWait:                 ackWait,
		MaxInflight:             maxInflight,
		SendWorkers:             sendWorkers,
		MaxDeliver:              maxDeliver,
	}
}

type Subscriber struct {
	stream       nats.JetStreamContext
	emailUseCase EmailUseCase
	logger       logging.Logger
	cfg          SubscriberCfg
}

func NewSubscriber(
	stream nats.JetStreamContext,
	emailUseCase EmailUseCase,
	subscriberCfg SubscriberCfg,
	logger logging.Logger) *Subscriber {

	return &Subscriber{
		stream:       stream,
		cfg:          subscriberCfg,
		emailUseCase: emailUseCase,
		logger:       logger,
	}

}

func (s *Subscriber) Subscribe(
	subject string,
	qgroup string,
	durableName string,
	maxDeliver int,
	ackWait time.Duration,
	maxInflight int,
	workersNum int,
	handler nats.MsgHandler) {
	s.logger.Infof("Subscribing to Subject: %v, group: %v", subject, qgroup)
	wg := &sync.WaitGroup{}
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go s.runWorker(
			wg,
			i,
			subject,
			qgroup,
			handler,
			nats.Durable(durableName),
			nats.MaxDeliver(maxDeliver),
			nats.AckWait(ackWait),
			nats.DeliverAll(),
			nats.MaxAckPending(maxInflight),
		)
	}
	wg.Wait()
}

func (s *Subscriber) Run(ctx context.Context) {
	go s.Subscribe(
		s.cfg.SendEmailSubject,
		s.cfg.EmailGroupName,
		s.cfg.DurableName,
		s.cfg.MaxDeliver,
		s.cfg.AckWait,
		s.cfg.MaxInflight,
		s.cfg.SendWorkers,
		s.processSendEmail(ctx),
	)
}

func (s *Subscriber) runWorker(
	wg *sync.WaitGroup,
	workerID int,
	subject string,
	qgroup string,
	handler nats.MsgHandler,
	opts ...nats.SubOpt,
) {

	s.logger.Infof("Subscribing worker: %v, subject: %v, qgroup: %v", workerID, subject, qgroup)
	defer wg.Done()
	sub, err := s.stream.QueueSubscribe(subject, qgroup, handler, opts...)
	if err != nil {
		s.logger.Errorf("WorkerID: %v, QueueSubscribe: %v", workerID, err)
		if err := sub.Unsubscribe(); err != nil {
			s.logger.Errorf("WorkerID: %v, sub.Unsubscribe error: %v", workerID, err)
		}
	}
}

func (s *Subscriber) processSendEmail(ctx context.Context) nats.MsgHandler {
	return func(msg *nats.Msg) {
		s.logger.Infof("subscriber process Send Email: %s", msg.Subject)
		var m model.Email
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			s.logger.Errorf("json.Unmarshal : %v", err)
			return
		}
		if err := retry.Do(func() error {
			return s.emailUseCase.Send(ctx, &m)
		},
			retry.Attempts(retryAttempts),
			retry.Delay(retryDelay),
			retry.Context(ctx),
		); err != nil {

			if err := s.publishErrorMessage(ctx, msg, err); err != nil {
				s.logger.Errorf("publishErrorMessage : %v", err)
				return
			}

			if err := msg.Ack(); err != nil {
				s.logger.Errorf("msg.Ack: %v", err)
				return
			}
			return
		}
		if err := msg.Ack(); err != nil {
			s.logger.Errorf("msg.Ack: %v", err)
			return
		}
	}
}

func (s *Subscriber) publishErrorMessage(ctx context.Context, msg *nats.Msg, err error) error {
	s.logger.Infof("publish dead letter queue message: %v", msg)
	errMsg := model.EmailErrorMsg{
		Subject: msg.Subject,
		Reply:   msg.Reply,
		Data:    msg.Data,
		Error:   err,
		Time:    time.Now().UTC(),
	}

	errMsgBytes, err := json.Marshal(&errMsg)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}
	_, err = s.stream.Publish(s.cfg.DeadMessageQueueSubject, errMsgBytes)
	return err
}

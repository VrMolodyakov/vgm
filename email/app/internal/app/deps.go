package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/email"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/interceptor"
	grpcMetrics "github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/metrics"
	httpMetrics "github.com/VrMolodyakov/vgm/email/app/internal/controller/http/v1/metrics"
	jet "github.com/VrMolodyakov/vgm/email/app/internal/controller/nats"
	natsMetrics "github.com/VrMolodyakov/vgm/email/app/internal/controller/nats/metrics"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/usecase"
	"github.com/VrMolodyakov/vgm/email/app/pkg/client/gmail"
	natsStream "github.com/VrMolodyakov/vgm/email/app/pkg/client/nats"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Deps struct {
	subscriber    *jet.Subscriber
	server        *grpc.Server
	metricsServer *http.Server
	connection    *nats.Conn
}

func (d *Deps) Setup(cfg *config.Config, logger logging.Logger) error {
	var conn *nats.Conn
	conn, streamCtx, err := natsStream.NewStreamContext(
		cfg.Nats.Host,
		cfg.Nats.Port,
		cfg.Subscriber.MainSubjectName,
		cfg.Subscriber.MainSubjects,
	)
	if err != nil {
		return err
	}
	d.connection = conn

	pub := jet.NewPublisher(streamCtx)
	emailClient := gmail.NewMailClient(
		cfg.Mail.SmtpAuthAddress,
		cfg.Mail.SmtpServerAddress,
		cfg.Mail.Name,
		cfg.Mail.FromAddress,
		cfg.Mail.FromPassword,
	)
	emailUseCase := usecase.NewEmailUseCase(logger, pub, cfg.Subscriber.SendEmailSubject, emailClient)
	subCfg := d.ReadSubscriberCfg(cfg.Subscriber)
	d.subscriber = jet.NewSubscriber(streamCtx, emailUseCase, subCfg, logger)

	natsMetrics.RegisterNatsMetrics()
	grpcMetrics.RegisterGrpcMetrics()
	d.metricsServer = httpMetrics.NewServer(cfg.MetricsServer)

	emailServer := email.NewServer(emailUseCase, logger, emailPb.UnimplementedEmailServiceServer{})
	serverOptions := []grpc.ServerOption{}
	if enableTLS {
		tlsCredentials, err := d.loadTLSCredentials()
		if err != nil {
			return err
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}
	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(
		interceptor.NewLoggerInterceptor(logger),
	))
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	d.server = grpc.NewServer(serverOptions...)

	emailPb.RegisterEmailServiceServer(d.server, emailServer)
	return nil
}

func (d *Deps) ReadSubscriberCfg(cfg config.Subscriber) jet.SubscriberCfg {
	return jet.NewSubscriberCfg(
		cfg.DurableName,
		cfg.DeadMessageSubject,
		cfg.SendEmailSubject,
		cfg.EmailGroupName,
		time.Duration(cfg.AckWait)*time.Second,
		cfg.Workers,
		cfg.MaxInflight,
		cfg.MaxDeliver,
	)
}

func (d *Deps) Close(ctx context.Context, logger logging.Logger) {
	if d.server != nil {
		d.server.Stop()
	}

	if d.metricsServer != nil {
		if err := d.metricsServer.Shutdown(ctx); err != nil {
			logger.Error(err, "shutdown metrics server")
		}
	}

	if d.connection != nil {
		d.connection.Close()
	}
}

func (d *Deps) loadTLSCredentials() (credentials.TransportCredentials, error) {
	dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
	path := containerConfigPath + clientCACertFile
	pemClientCA, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair(containerConfigPath+serverCertFile, containerConfigPath+serverKeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

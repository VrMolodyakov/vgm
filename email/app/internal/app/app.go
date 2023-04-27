package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"net"

	"github.com/VrMolodyakov/vgm/email/app/internal/config"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/email"
	"github.com/VrMolodyakov/vgm/email/app/internal/controller/grpc/v1/interceptor"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	emailPb "github.com/VrMolodyakov/vgm/email/app/gen/go/proto/email/v1"
	jet "github.com/VrMolodyakov/vgm/email/app/internal/controller/nats"
	"github.com/VrMolodyakov/vgm/email/app/internal/domain/email/usecase"
	"github.com/VrMolodyakov/vgm/email/app/pkg/client/gmail"
	"github.com/VrMolodyakov/vgm/email/app/pkg/client/nats"
	"github.com/VrMolodyakov/vgm/email/app/pkg/jaeger"
	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

//TODO:put in yaml
const (
	enableTLS               = true
	serverCertFile   string = "cert/email-server-cert.pem"
	serverKeyFile    string = "cert/email-server-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
	serviceName      string = "email-server"
)

type app struct {
	cfg        *config.Config
	logger     logging.Logger
	grpcServer *grpc.Server
}

func NewApp(cfg *config.Config, logger logging.Logger) *app {
	return &app{cfg: cfg, logger: logger}
}

func (a *app) Run(ctx context.Context) {
	streamCtx := nats.NewStreamContext(
		a.cfg.Nats.Host,
		a.cfg.Nats.Port,
		a.cfg.Subscriber.MainSubjectName,
		a.cfg.Subscriber.MainSubjects,
	)
	info, err := streamCtx.StreamInfo("email")
	if err != nil {
		a.logger.Fatal(err)
	}
	a.logger.Info("cluster", info.Cluster)
	a.logger.Info("config", info.Config)
	a.logger.Info("config", info.Sources)
	pub := jet.NewPublisher(streamCtx)
	emailClient := gmail.NewMailClient(
		a.cfg.Mail.SmtpAuthAddress,
		a.cfg.Mail.SmtpServerAddress,
		a.cfg.Mail.Name,
		a.cfg.Mail.FromAddress,
		a.cfg.Mail.FromPassword,
	)
	emailUseCase := usecase.NewEmailUseCase(a.logger, pub, a.cfg.Subscriber.SendEmailSubject, emailClient)
	go func() {
		subCfg := jet.NewSubscriberCfg(
			a.cfg.Subscriber.DurableName,
			a.cfg.Subscriber.DeadMessageSubject,
			a.cfg.Subscriber.SendEmailSubject,
			a.cfg.Subscriber.EmailGroupName,
			time.Duration(a.cfg.Subscriber.AckWait)*time.Second,
			a.cfg.Subscriber.Workers,
			a.cfg.Subscriber.MaxInflight,
			a.cfg.Subscriber.MaxDeliver,
		)
		sub := jet.NewSubscriber(streamCtx, emailUseCase, subCfg, a.logger)
		sub.Run(ctx)
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	fmt.Printf("print ip = %s port = %d", a.cfg.GRPC.IP, a.cfg.GRPC.Port)
	a.logger.Info("grpc listener :=", zap.String("ip", a.cfg.GRPC.IP), zap.Int("port", a.cfg.GRPC.Port))
	if err != nil {
		a.logger.Error(err.Error())
	}

	err = jaeger.SetGlobalTracer(serviceName, a.cfg.Jaeger.Address, a.cfg.Jaeger.Port)
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	emailServer := email.NewServer(emailUseCase, a.logger, emailPb.UnimplementedEmailServiceServer{})
	serverOptions := []grpc.ServerOption{}
	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			a.logger.Fatalf("cannot load TLS credentials: %s", err.Error())
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}
	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(
		interceptor.NewLoggerInterceptor(a.logger),
	))
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	a.grpcServer = grpc.NewServer(serverOptions...)

	emailPb.RegisterEmailServiceServer(a.grpcServer, emailServer)
	reflection.Register(a.grpcServer)
	a.logger.Info("start grpc serve")
	a.grpcServer.Serve(listener)
	a.logger.Info("end of email service")
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
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

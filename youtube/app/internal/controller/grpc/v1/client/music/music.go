package music

import (
	"context"
	"log"

	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/youtube/internal/controller/grpc/v1/client"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tracer = otel.Tracer("music-grpc-client")
)

type musicClient struct {
	target string
	client albumPb.MusicServiceClient
}

func NewMusicClient(target string) *musicClient {
	if target == "" {
		log.Fatalln("Error in Access to GRPC URL in music client")
	}
	return &musicClient{
		target: target,
	}
}

func (m *musicClient) Start() {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(m.target, transportOption)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	m.client = albumPb.NewMusicServiceClient(conn)
}

func (m *musicClient) StartWithTLS(certs client.ClientCerts) {

	tlsCredentials, err := certs.LoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	transportOption := grpc.WithTransportCredentials(tlsCredentials)

	conn, err := grpc.Dial(
		m.target,
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		transportOption)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	m.client = albumPb.NewMusicServiceClient(conn)
}

func (m *musicClient) FindRandomTitles(ctx context.Context, count uint64) ([]string, error) {
	ctx, span := tracer.Start(ctx, "client.FindRandomTitles")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	request := albumPb.FindRandomTitlesRequest{
		Count: count,
	}
	pb, err := m.client.FindRandomTitles(ctx, &request)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return pb.GetTitles(), nil
}

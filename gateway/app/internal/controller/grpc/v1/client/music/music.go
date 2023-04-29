package music

import (
	"context"
	"log"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/grpc/v1/client"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
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

func (m *musicClient) StartWithTSL(certs client.ClientCerts) {

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

func (m *musicClient) CreateAlbum(ctx context.Context, album model.Album) error {
	_, span := tracer.Start(ctx, "client.CreateAlbum")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	tracklist := make([]*albumPb.Track, len(album.Tracklist))
	for i := 0; i < len(album.Tracklist); i++ {
		tracklist[i] = album.Tracklist[i].PbFromkModel()
	}
	credits := make([]*albumPb.Credit, len(album.Credits))
	for i := 0; i < len(album.Credits); i++ {
		credits[i] = album.Credits[i].PbFromkModel()
	}

	request := albumPb.CreateAlbumRequest{
		Title:          album.Album.Title,
		ReleasedAt:     album.Album.ReleasedAt,
		CatalogNumber:  album.Info.CatalogNumber,
		FullImageSrc:   &album.Info.FullImageSrc,
		SmallImageSrc:  &album.Info.SmallImageSrc,
		Barcode:        &album.Info.Barcode,
		Price:          album.Info.Price,
		CurrencyCode:   album.Info.CurrencyCode,
		MediaFormat:    album.Info.MediaFormat,
		Classification: album.Info.Classification,
		Publisher:      album.Info.Publisher,
		Tracklist:      tracklist,
		Credits:        credits,
	}

	_, err := m.client.CreateAlbum(ctx, &request)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (m *musicClient) CreatePerson(ctx context.Context, person model.Person) error {
	_, span := tracer.Start(ctx, "client.CreatePerson")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	request := albumPb.CreatePersonRequest{
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: person.BirthDate,
	}
	_, err := m.client.CreatePerson(ctx, &request)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (m *musicClient) FindAll(
	ctx context.Context,
	pagination model.Pagination,
	titleView model.AlbumTitleView,
	releaseView model.AlbumReleasedView,
	sort model.Sort,
) ([]model.AlbumView, error) {

	_, span := tracer.Start(ctx, "client.FindAll")
	defer span.End()

	request := albumPb.FindAllAlbumsRequest{
		Pagination: pagination.PbFromModel(),
		Title:      titleView.PbFromModel(),
		ReleasedAt: releaseView.PbFromModel(),
		Sort:       sort.PBFromModel(),
	}

	pb, err := m.client.FindAllAlbums(ctx, &request)
	if err != nil {
		return nil, err
	}
	albumsPb := pb.GetAlbums()
	albums := make([]model.AlbumView, len(albumsPb))
	for i := 0; i < len(albums); i++ {
		albums[i] = model.AlbumFromPb(albumsPb[i])
	}
	return albums, nil
}

func (m *musicClient) FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error) {
	_, span := tracer.Start(ctx, "client.FindFullAlbum")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	request := albumPb.FindFullAlbumRequest{
		AlbumId: id,
	}
	pb, err := m.client.FindFullAlbum(ctx, &request)
	if err != nil {
		logger.Error(err.Error())
		return model.FullAlbum{}, err
	}
	return model.FullAlbumFromPb(pb), nil
}

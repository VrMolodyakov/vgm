package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

//TODO:put in yaml
const (
	enableTLS               = true
	clientCertFile   string = "cert/client-cert.pem"
	clientKeyFile    string = "cert/client-key.pem"
	clientCACertFile string = "cert/ca-cert.pem"
)

type musicClient struct {
	target string
	client albumPb.AlbumServiceClient
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
	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	conn, err := grpc.Dial(m.target, transportOption)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	m.client = albumPb.NewAlbumServiceClient(conn)
}

func (m *musicClient) Create(ctx context.Context, album model.Album) error {
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
		ImageSrc:       &album.Info.ImageSrc,
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

func (m *musicClient) FindAll(
	ctx context.Context,
	pagination model.Pagination,
	titleView model.AlbumTitleView,
	releaseView model.AlbumReleasedView,
	sort model.Sort,
) ([]model.AlbumView, error) {

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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
	path := containerConfigPath + clientCACertFile
	pemServerCA, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(containerConfigPath+clientCertFile, containerConfigPath+clientKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

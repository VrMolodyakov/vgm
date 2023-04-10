package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"google.golang.org/grpc/credentials"
)

type ClientCerts struct {
	EnableTLS        bool
	ClientCertFile   string
	ClientKeyFile    string
	ClientCACertFile string
}

func NewClientCerts(
	enableTLS bool,
	clientCertFile string,
	clientKeyFile string,
	clientCACertFile string,
) ClientCerts {
	return ClientCerts{
		EnableTLS:        enableTLS,
		ClientCertFile:   clientCertFile,
		ClientKeyFile:    clientKeyFile,
		ClientCACertFile: clientCACertFile,
	}
}

func (c *ClientCerts) loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
	path := containerConfigPath + c.ClientCACertFile
	pemServerCA, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(
		containerConfigPath+c.ClientCertFile,
		containerConfigPath+c.ClientKeyFile,
	)
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

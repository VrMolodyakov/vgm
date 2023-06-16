package config

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var instance *Config
var once sync.Once

type Youtube struct {
	ApiKey string `env:"YOUTUBE_API_KEY"       env-required:""`
}

type MusicGRPC struct {
	HostName string `env:"MUSIC_SERVICE_NAME"`
	Port     int    `env:"MUSIC_GRPC_PORT"`
}

type CORS struct {
	AllowedMethods     []string `yaml:"allowed_methods"`
	AllowedOrigins     []string `yaml:"allowed_origins"`
	AllowCredentials   bool     `yaml:"allow_credentials"`
	AllowedHeaders     []string `yaml:"allowed_headers"`
	OptionsPassthrough bool     `yaml:"options_passthrough"`
	ExposedHeaders     []string `yaml:"exposed_headers"`
	Debug              bool     `yaml:"debug"`
}

type YoutubeClientCert struct {
	EnableTLS        bool   `yaml:"enable_tls"`
	ClientCertFile   string `yaml:"yt-client_cert_file"`
	ClientKeyFile    string `yaml:"yt-client_key_file"`
	ClientCACertFile string `yaml:"client_CAcert_file"`
}

type Logger struct {
	DisableCaller     bool   `yaml:"disable_caller"`
	Development       bool   `yaml:"development"`
	DisableStacktrace bool   `yaml:"disable_stacktrace"`
	Encoding          string `yaml:"encoding"`
	Level             string `yaml:"level"`
}

type Jaeger struct {
	ServiceName string `env:"YOUTUBE_SERVICE_NAME"`
	Address     string `env:"JAEGER_ADDRESS"`
	Port        string `env:"JAEGER_PORT"`
}

type YoutubeServer struct {
	IP           string `env:"YOUTUBE_GRPC_IP"`
	Port         int    `env:"YOUTUBE_GRPC_PORT"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type Config struct {
	Logger            Logger            `yaml:"logger"`
	YoutubeClientCert YoutubeClientCert `yaml:"youtube_client"`
	CORS              CORS              `yaml:"cors"`
	YoutubeServer     YoutubeServer     `yaml:"youtube_server"`
	Youtube           Youtube
	Jaeger            Jaeger
	MusicGRPC         MusicGRPC
}

func GetConfig() (*Config, error) {
	var cfgErr error
	once.Do(func() {
		instance = &Config{}
		dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
		if exist, _ := Exists(containerConfigPath + "/configs/config.yaml"); exist {
			if err := cleanenv.ReadConfig(containerConfigPath+"/configs/config.yaml", instance); err != nil {
				cfgErr = err
			}
		}
	})
	return instance, cfgErr
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

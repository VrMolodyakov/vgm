package config

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Redis struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DbNumber int    `env:"REDIS_DBNUMBER"`
}

type KeyPairs struct {
	AccessPublic   string `env:"ACCESS_PUBLIC"`
	AccessPrivate  string `env:"ACCESS_PRIVATE"`
	RefreshPublic  string `env:"REFRESH_PUBLIC"`
	RefreshPrivate string `env:"REFRESH_PRIVATE"`
	AccessTtl      int    `env:"ACCESS_TTL"`
	RefreshTtl     int    `env:"REFRESH_TTL"`
}

type Postgres struct {
	User     string `env:"USERDB_POSTGRES_USER"     env-required:""`
	Password string `env:"USERDB_POSTGRES_PASSWORD" env-required:""`
	Database string `env:"USERDB_POSTGRES_DB"       env-required:""`
	IP       string `env:"USERDB_POSTGRES_IP"       env-required:""`
	Port     string `env:"USERDB_POSTGRES_PORT"       env-required:""`
	PoolSize string `env:"USERDB_POSTGRES_POOL_SIZE"       env-required:""`
}

type UserServer struct {
	IP           string `env:"USER_SERVER_IP"`
	Port         int    `env:"USER_SERVER_PORT"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type MusicServer struct {
	IP           string `env:"MUSIC_SERVER_IP"`
	Port         int    `env:"MUSIC_SERVER_PORT"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
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

type MusicGRPC struct {
	HostName string `env:"MUSIC_SERVICE_NAME"`
	Port     int    `env:"MUSIC_GRPC_PORT"`
}

type EmailGRPC struct {
	HostName string `env:"EMAIL_SERVICE_NAME"`
	Port     int    `env:"EMAIL_GRPC_PORT"`
}

type EmailClientCert struct {
	EnableTLS        bool   `yaml:"enable_tls"`
	ClientCertFile   string `yaml:"client_cert_file"`
	ClientKeyFile    string `yaml:"client_key_file"`
	ClientCACertFile string `yaml:"client_CAcert_file"`
}

type MusicClientCert struct {
	EnableTLS        bool   `yaml:"enable_tls"`
	ClientCertFile   string `yaml:"client_cert_file"`
	ClientKeyFile    string `yaml:"client_key_file"`
	ClientCACertFile string `yaml:"client_CAcert_file"`
}

type Jaeger struct {
	ServiceName string `env:"GATEWAY_SERVICE_NAME"`
	Address     string `env:"JAEGER_ADDRESS"`
	Port        string `env:"JAEGER_PORT"`
}

type MetricsServer struct {
	Port         int    `yaml:"port"`
	IP           string `yaml:"ip"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type Config struct {
	CORS            CORS          `yaml:"cors"`
	MusicServer     MusicServer   `yaml:"music_server"`
	UserServer      UserServer    `yaml:"user_server"`
	MetricsServer   MetricsServer `yaml:"metrics_server"`
	Postgres        Postgres
	MusicGRPC       MusicGRPC
	EmailGRPC       EmailGRPC
	Redis           Redis
	KeyPairs        KeyPairs
	Jaeger          Jaeger
	EmailClientCert EmailClientCert `yaml:"email_client"`
	MusicClientCert MusicClientCert `yaml:"music_client"`
}

var instance *Config
var once sync.Once

//TODO: put config here?
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

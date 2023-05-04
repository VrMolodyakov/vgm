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

type Postgres struct {
	User     string `env:"POSTGRES_USER"     env-required:""`
	Password string `env:"POSTGRES_PASSWORD" env-required:""`
	Database string `env:"POSTGRES_DB"       env-required:""`
	IP       string `env:"POSTGRES_IP"       env-required:""`
	Port     string `env:"POSTGRES_PORT"       env-required:""`
	PoolSize string `env:"POSTGRES_POOL_SIZE"       env-required:""`
}

type Jaeger struct {
	ServiceName string `env:"MUSIC_SERVICE_NAME"`
	Address     string `env:"JAEGER_ADDRESS"`
	Port        string `env:"JAEGER_PORT"`
}

type GRPC struct {
	IP   string `env:"MUSIC_GRPC_IP"`
	Port int    `env:"MUSIC_GRPC_PORT"`
}

type Config struct {
	Postgres Postgres
	Jaeger   Jaeger
	GRPC     GRPC
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

package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "\\configs\\config.yaml"
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

type GRPC struct {
	IP   string `env:"MUSIC_GRPC_IP"`
	Port int    `env:"MUSIC_GRPC_PORT"`
}

type Config struct {
	Postgres Postgres
	GRPC     GRPC
}

//TODO: remote root path
func GetConfig() *Config {
	once.Do(func() {
		rootPath, _ := os.Getwd()
		root := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(rootPath))))
		instance = &Config{}
		path := root + configPath
		dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		containerConfigPath := filepath.Dir(filepath.Dir(dockerPath))
		fmt.Println("container docker path : ", containerConfigPath)
		if exist, _ := Exists(path); exist {
			if err := cleanenv.ReadConfig(path, instance); err != nil {
				log.Fatal(err)
			}
		} else if exist, _ := Exists(containerConfigPath + "/configs/config.yaml"); exist {
			fmt.Println("inside docker path")
			if err := cleanenv.ReadConfig(containerConfigPath+"/configs/config.yaml", instance); err != nil {
				log.Fatal(err)
			}
		}
	})
	return instance
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

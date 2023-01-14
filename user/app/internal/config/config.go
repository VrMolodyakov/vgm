package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "\\configs\\config.yaml"
)

type Postgres struct {
	User     string `env:"POSTGRES_USER"     env-required:""`
	Password string `env:"POSTGRES_PASSWORD" env-required:""`
	Database string `env:"POSTGRES_IP"       env-required:""`
}

type HTTP struct {
	IP           string        `yaml:"ip"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
	CORS         struct {
		AllowedMethods     []string `yaml:"allowed_methods"`
		AllowedOrigins     []string `yaml:"allowed_origins"`
		AllowCredentials   bool     `yaml:"allow_credentials"`
		AllowedHeaders     []string `yaml:"allowed_headers"`
		OptionsPassthrough bool     `yaml:"options_passthrough"`
		ExposedHeaders     []string `yaml:"exposed_headers"`
		Debug              bool     `yaml:"debug"`
	} `yaml:"cors"`
}

type Config struct {
	HTTP     HTTP `yaml:"http"`
	Postgres Postgres
}

var instance *Config
var once sync.Once

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

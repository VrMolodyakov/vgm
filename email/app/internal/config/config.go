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

type Mail struct {
	SmtpAuthAddress   string `env:"SMTP_AUTH_ADDRESS"`
	SmtpServerAddress string `env:"SMTP_SERVER_ADDRESS"`
	Name              string `env:"EMAIL_NAME"`
	FromAddress       string `env:"EMAIL_SENDER_ADDRESS"`
	FromPassword      string `env:"EMAIL_SENDER_PASSWORD"`
}

type Subscriber struct {
	DurableName        string `yaml:"durable_name"`
	DeadMessageSubject string `yaml:"dead_message_subject"`
	SendEmailSubject   string `yaml:"send_subject"`
	EmailGroupName     string `yaml:"email_group_name"`
	MainSubject        string `yaml:"main_subject"`
	AckWait            int    `yaml:"ack_wait"`
	MaxInflight        int    `yaml:"max_inflight"`
	MaxDeliver         int    `yaml:"max_deliver"`
	Workers            int    `yaml:"workers"`
}

type GRPC struct {
	IP   string `env:"EMAIL_GRPC_IP"`
	Port int    `env:"EMAIL_GRPC_PORT"`
}

type Nats struct {
	Host string `env:"NATS_HOST"`
	Port int    `env:"NATS_PORT"`
}

type Logger struct {
	DisableCaller     bool   `yaml:"disable_caller"`
	Development       bool   `yaml:"development"`
	DisableStacktrace bool   `yaml:"disable_stacktrace"`
	Encoding          string `yaml:"encoding"`
	Level             string `yaml:"level"`
}

type Config struct {
	Logger     Logger     `yaml:"logger"`
	Subscriber Subscriber `yaml:"subscriber"`
	Mail       Mail
	Nats       Nats
	GRPC       GRPC
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

package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port   string `yaml : "port"`
	Host   string `yaml : "host"`
	LogLvl string `yaml : "loglvl"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		path, _ := os.Getwd()
		root := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(path))))
		if err := cleanenv.ReadConfig(root+"\\configs\\config.yaml", instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}

package config

import "time"

type Mail struct {
	From           string        `yaml:"from"`
	Host           string        `yaml:"host"`
	Port           int           `env:"MAIL_PORT"`
	Username       string        `yaml:"username"`
	Password       string        `yaml:"password"`
	KeepAlive      bool          `yaml:"keep_alive"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	SendTimeout    time.Duration `yaml:"send_timeout"`
}

type Config struct {
	Mail Mail `yaml:"mail"`
}

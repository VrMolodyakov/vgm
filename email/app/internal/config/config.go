package config

type Mail struct {
	SmtpAuthAddress   string `env:"SMTP_AUTH_ADDRESS"`
	SmtpServerAddress string `env:"SMTP_SERVER_ADDRESS"`
	Name              string `env:"EMAIL_NAME"`
	FromAddress       string `env:"EMAIL_SENDER_ADDRESS"`
	FromPassword      string `env:"EMAIL_SENDER_PASSWORD"`
}

type GRPC struct {
	IP   string `env:"EMAIL_GRPC_IP"`
	Port int    `env:"EMAIL_GRPC_PORT"`
}

type Nats struct {
	Host string `env:"NATS_PORT"`
	Port int    `env:"NATS_PORT"`
}

type Config struct {
	Mail Mail
	Nats Nats
}

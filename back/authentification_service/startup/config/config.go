package config

import "os"

type Config struct {
	Port                       string
	AuthentificationDBHost     string
	AuthentificationDBPort     string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	RegisterUserCommandSubject string
	RegisterUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                       os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		AuthentificationDBHost:     os.Getenv("AUTHENTIFICATION_DB_HOST"),
		AuthentificationDBPort:     os.Getenv("AUTHENTIFICATION_DB_PORT"),
		NatsHost:                   os.Getenv("NATS_HOST"),
		NatsPort:                   os.Getenv("NATS_PORT"),
		NatsUser:                   os.Getenv("NATS_USER"),
		NatsPass:                   os.Getenv("NATS_PASS"),
		RegisterUserCommandSubject: os.Getenv("REGISTER_USER_COMMAND_SUBJECT"),
		RegisterUserReplySubject:   os.Getenv("REGISTER_USER_REPLY_SUBJECT"),
	}
}

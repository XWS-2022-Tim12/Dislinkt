package config

import "os"

type Config struct {
	Port       string
	MessageDBHost string
	MessageDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost: os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort: os.Getenv("MESSAGE_DB_PORT"),
	}
}

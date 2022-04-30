package config

import "os"

type Config struct {
	Port       string
	AuthentificationDBHost string
	AuthentificationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		AuthentificationDBHost: os.Getenv("AUTHENTIFICATION_DB_HOST"),
		AuthentificationDBPort: os.Getenv("AUTHENTIFICATION_DB_PORT"),
	}
}

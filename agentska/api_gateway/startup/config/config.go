package config

import "os"

type Config struct {
	Port                 string
	UserHost             string
	UserPort             string
	AuthentificationHost string
	AuthentificationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("GATEWAY_PORT"),
		UserHost:             os.Getenv("USER_SERVICE_HOST"),
		UserPort:             os.Getenv("USER_SERVICE_PORT"),
		AuthentificationHost: os.Getenv("AUTHENTIFICATION_SERVICE_HOST"),
		AuthentificationPort: os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
	}
}

package config

import "os"

type Config struct {
	Port                 string
	UserHost             string
	UserPort             string
	AuthentificationHost string
	AuthentificationPort string
	PostHost             string
	PostPort             string
}

func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("GATEWAY_PORT"),
		UserHost:             os.Getenv("USER_SERVICE_HOST"),
		UserPort:             os.Getenv("USER_SERVICE_PORT"),
		AuthentificationHost: os.Getenv("AUTHENTIFICATION_SERVICE_HOST"),
		AuthentificationPort: os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		PostHost:             os.Getenv("POST_SERVICE_HOST"),
		PostPort:             os.Getenv("POST_SERVICE_PORT"),
	}
}

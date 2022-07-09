package config

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	Uri      string
	Username string
	Password string
}

func NewConfig() *Config {

	return &Config{
		Host:     os.Getenv("USER_SUGGESTIONS_SERVICE_HOST"),
		Port:     os.Getenv("USER_SUGGESTIONS_SERVICE_PORT"),
		Uri:      "neo4j://neo4j:7687",
		Username: os.Getenv("USER_SUGGESTIONS_DB_USER"),
		Password: os.Getenv("USER_SUGGESTIONS_DB_PASS"),
	}
}

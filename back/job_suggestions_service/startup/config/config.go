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
		Host:     os.Getenv("JOB_SUGGESTIONS_SERVICE_HOST"),
		Port:     os.Getenv("JOB_SUGGESTIONS_SERVICE_PORT"),
		Uri:      "neo4j://neo4j:7687",
		Username: os.Getenv("JOB_SUGGESTIONS_DB_USER"),
		Password: os.Getenv("JOB_SUGGESTIONS_DB_PASS"),
	}
}

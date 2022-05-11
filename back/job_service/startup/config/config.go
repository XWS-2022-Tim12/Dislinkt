package config

import "os"

type Config struct {
	Port      string
	JobDBHost string
	JobDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:      os.Getenv("JOB_SERVICE_PORT"),
		JobDBHost: os.Getenv("JOB_DB_HOST"),
		JobDBPort: os.Getenv("JOB_DB_PORT"),
	}
}

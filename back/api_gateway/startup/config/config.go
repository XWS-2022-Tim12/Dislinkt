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
	MessageHost          string
	MessagePort          string
	JobHost              string
	JobPort              string
	JobSuggestionsHost   string
	JobSuggestionsPort   string
	UserSuggestionsHost  string
	UserSuggestionsPort  string
	NotificationHost     string
	NotificationPort     string
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
		MessageHost:          os.Getenv("MESSAGE_SERVICE_HOST"),
		MessagePort:          os.Getenv("MESSAGE_SERVICE_PORT"),
		JobHost:              os.Getenv("JOB_SERVICE_HOST"),
		JobPort:              os.Getenv("JOB_SERVICE_PORT"),
		JobSuggestionsHost:   os.Getenv("JOB_SUGGESTIONS_SERVICE_HOST"),
		JobSuggestionsPort:   os.Getenv("JOB_SUGGESTIONS_SERVICE_PORT"),
		UserSuggestionsHost:  os.Getenv("USER_SUGGESTIONS_SERVICE_HOST"),
		UserSuggestionsPort:  os.Getenv("USER_SUGGESTIONS_SERVICE_PORT"),
		NotificationHost:     os.Getenv("NOTIFICATION_SERVICE_HOST"),
		NotificationPort:     os.Getenv("NOTIFICATION_SERVICE_PORT"),
	}
}

package main

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

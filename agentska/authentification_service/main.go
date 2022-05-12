package main

import (
	"github.com/XWS-2022-Tim12/Dislinkt/agentska/authentification_service/startup"
	cfg "github.com/XWS-2022-Tim12/Dislinkt/agentska/authentification_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

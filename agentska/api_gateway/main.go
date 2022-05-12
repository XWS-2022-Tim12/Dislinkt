package main

import (
	"github.com/XWS-2022-Tim12/Dislinkt/agentska/api_gateway/startup"
	"github.com/XWS-2022-Tim12/Dislinkt/agentska/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

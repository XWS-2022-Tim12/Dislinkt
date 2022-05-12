package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/infrastructure/api"
	cfg "github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/startup/config"
	authentificationGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	jobGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	postGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	userGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	authentificationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthentificationHost, server.config.AuthentificationPort)
	err = authentificationGw.RegisterAuthentificationServiceHandlerFromEndpoint(context.TODO(), server.mux, authentificationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if err != nil {
		panic(err)
	}

	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	err = jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	authentificationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthentificationHost, server.config.AuthentificationPort)
	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	authentificationHandler := api.NewAuthentificationHandler(authentificationEndpoint, userEndpoint, postEndpoint, jobEndpoint)
	authentificationHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}

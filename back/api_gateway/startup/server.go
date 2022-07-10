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
	jobSuggestionsGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	postGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	messageGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/message_service"
	userGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	userSuggestionsGw "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
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

	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	err = messageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, messageEndpoint, opts)
	if err != nil {
		panic(err)
	}

	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	err = jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEndpoint, opts)
	if err != nil {
		panic(err)
	}
	jobSuggestionsEndpoint := fmt.Sprintf("%s:%s", server.config.JobSuggestionsHost, server.config.JobSuggestionsPort)
	err = jobSuggestionsGw.RegisterJobSuggestionsServiceHandlerFromEndpoint(context.TODO(), server.mux, jobSuggestionsEndpoint, opts)
	if err != nil {
		panic(err)
	}
	userSuggestionsEndpoint := fmt.Sprintf("%s:%s", server.config.UserSuggestionsHost, server.config.UserSuggestionsPort)
	err = userSuggestionsGw.RegisterUserSuggestionsServiceHandlerFromEndpoint(context.TODO(), server.mux, userSuggestionsEndpoint, opts)

	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	authentificationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthentificationHost, server.config.AuthentificationPort)
	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	jobSuggestionsEndpoint := fmt.Sprintf("%s:%s", server.config.JobSuggestionsHost, server.config.JobSuggestionsPort)
	userSuggestionsEndpoint := fmt.Sprintf("%s:%s", server.config.UserSuggestionsHost, server.config.UserSuggestionsPort)
	authentificationHandler := api.NewAuthentificationHandler(authentificationEndpoint, userEndpoint, postEndpoint, jobEndpoint, userSuggestionsEndpoint, jobSuggestionsEndpoint, messageEndpoint)
	authentificationHandler.Init(server.mux)
}

func (server *Server) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(server.mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}

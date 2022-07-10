package startup

import (
	"fmt"
	"log"
	"net"

	user_suggestions_service "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/startup/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "user_suggestions_service"
)

func (server *Server) Start() {
	neo4jsession := server.initNeo4jSession()
	userSuggestionsStore := server.initUserSuggestionsStore(neo4jsession)
	userSuggestionsService := server.initUserSuggestionsService(userSuggestionsStore)
	userSuggestionsHandler := server.initUserSuggestionsHandler(userSuggestionsService)

	server.startGrpcServer(userSuggestionsHandler)
}

func (server *Server) initNeo4jSession() *neo4j.Session {
	session, err := persistence.GetClient(server.config.Username, server.config.Password, server.config.Uri)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func (server *Server) initUserSuggestionsStore(client *neo4j.Session) domain.UserSuggestionsGraph {
	store := persistence.NewUserSuggestionsGraph(client)
	store.DeleteAll()
	return store
}

func (server *Server) initUserSuggestionsService(store domain.UserSuggestionsGraph) *application.UserSuggestionsService {
	return application.NewUserSuggestionsService(store)
}

func (server *Server) initUserSuggestionsHandler(service *application.UserSuggestionsService) *api.UserSuggestionsHandler {
	return api.NewUserSuggestionsHandler(service)
}

func (server *Server) startGrpcServer(userSuggestionsHandler *api.UserSuggestionsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	user_suggestions_service.RegisterUserSuggestionsServiceServer(grpcServer, userSuggestionsHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

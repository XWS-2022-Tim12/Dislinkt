package startup

import (
	"fmt"
	"log"
	"net"

	job_suggestions_service "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup/config"
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
	QueueGroup = "job_suggestions_service"
)

func (server *Server) Start() {
	neo4jsession := server.initNeo4jSession()
	jobSuggestionsStore := server.initJobSuggestionsStore(neo4jsession)
	jobSuggestionsService := server.initJobSuggestionsService(jobSuggestionsStore)
	jobSuggestionsHandler := server.initJobSuggestionsHandler(jobSuggestionsService)

	server.startGrpcServer(jobSuggestionsHandler)
}

func (server *Server) initNeo4jSession() *neo4j.Session {
	session, err := persistence.GetClient(server.config.Username, server.config.Password, server.config.Uri)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func (server *Server) initJobSuggestionsStore(client *neo4j.Session) domain.JobSuggestionsGraph {
	store := persistence.NewJobSuggestionsGraph(client)
	store.DeleteAll()
	return store
}

func (server *Server) initJobSuggestionsService(store domain.JobSuggestionsGraph) *application.JobSuggestionsService {
	return application.NewJobSuggestionsService(store)
}

func (server *Server) initJobSuggestionsHandler(service *application.JobSuggestionsService) *api.JobSuggestionsHandler {
	return api.NewJobSuggestionsHandler(service)
}

func (server *Server) startGrpcServer(jobSuggestionsHandler *api.JobSuggestionsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	job_suggestions_service.RegisterJobSuggestionsServiceServer(grpcServer, jobSuggestionsHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

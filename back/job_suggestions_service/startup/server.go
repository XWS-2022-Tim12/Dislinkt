package startup

import (
	"fmt"
	"io"
	"log"
	"net"

	job_suggestions_service "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup/config"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/tracer"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("job-suggestions-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)),
	)

	job_suggestions_service.RegisterJobSuggestionsServiceServer(grpcServer, jobSuggestionsHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

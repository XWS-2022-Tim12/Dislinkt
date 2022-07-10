package startup

import (
	"fmt"
	"io"
	"log"
	"net"

	job "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/startup/config"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/tracer"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("job-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	jobStore := server.initJobStore(mongoClient)

	JobService := server.initJobService(jobStore)

	jobHandler := server.initJobHandler(JobService)

	server.startGrpcServer(jobHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.JobDBHost, server.config.JobDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initJobStore(client *mongo.Client) domain.JobStore {
	store := persistence.NewJobMongoDBStore(client)
	store.DeleteAll()
	for _, job := range jobs {
		_, err := store.Insert(job)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initJobService(store domain.JobStore) *application.JobService {
	return application.NewJobService(store)
}

func (server *Server) initJobHandler(service *application.JobService) *api.JobHandler {
	return api.NewJobHandler(service)
}

func (server *Server) startGrpcServer(jobHandler *api.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)),
	)
	job.RegisterJobServiceServer(grpcServer, jobHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/startup/config"
	session "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	saga "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/messaging"
	"github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

const (
	QueueGroup = "authentification_service"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	sessionStore := server.initSessionStore(mongoClient)

	sessionService := server.initSessionService(sessionStore)

	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)

	server.initAddUserHandler(sessionService, replyPublisher, commandSubscriber)

	sessionHandler := server.initSessionHandler(sessionService)

	server.startGrpcServer(sessionHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthentificationDBHost, server.config.AuthentificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initSessionStore(client *mongo.Client) domain.SessionStore {
	store := persistence.NewSessionMongoDBStore(client)
	store.DeleteAll()
	return store
}

func (server *Server) initSessionService(store domain.SessionStore) *application.SessionService {
	return application.NewSessionService(store)
}

func (server *Server) initSessionHandler(service *application.SessionService) *api.SessionHandler {
	return api.NewSessionHandler(service)
}

func (server *Server) initAddUserHandler(service *application.SessionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateAuthentificationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) startGrpcServer(sessionHandler *api.SessionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	session.RegisterAuthentificationServiceServer(grpcServer, sessionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

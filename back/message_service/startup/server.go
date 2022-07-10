package startup

import (
	"fmt"
	"log"
	"net"

	message "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/message_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	messageStore := server.initMessageStore(mongoClient)

	messageService := server.initMessageService(messageStore)

	messageHandler := server.initMessageHandler(messageService)

	server.startGrpcServer(messageHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	store.DeleteAll()
	for _, message := range messages {
		_, err := store.Insert(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initMessageService(store domain.MessageStore) *application.MessageService {
	return application.NewMessageService(store)
}

func (server *Server) initMessageHandler(service *application.MessageService) *api.MessageHandler {
	return api.NewMessageHandler(service)
}

func (server *Server) startGrpcServer(messageHandler *api.MessageHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	message.RegisterMessageServiceServer(grpcServer, messageHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

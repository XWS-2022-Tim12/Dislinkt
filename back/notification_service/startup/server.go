package startup

import (
	"fmt"
	"log"
	"net"

	notification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/notification_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/startup/config"
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
	notificationStore := server.initNotificationStore(mongoClient)

	NotificationService := server.initNotificationService(notificationStore)

	notificationHandler := server.initNotificationHandler(NotificationService)

	server.startGrpcServer(notificationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationDBHost, server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationStore(client *mongo.Client) domain.NotificationStore {
	store := persistence.NewNotificationMongoDBStore(client)
	store.DeleteAll()
	for _, notification := range notifications {
		_, err := store.Insert(notification)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initNotificationService(store domain.NotificationStore) *application.NotificationService {
	return application.NewNotificationService(store)
}

func (server *Server) initNotificationHandler(service *application.NotificationService) *api.NotificationHandler {
	return api.NewNotificationHandler(service)
}

func (server *Server) startGrpcServer(notificationHandler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

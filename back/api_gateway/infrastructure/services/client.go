package services

import (
	"log"

	authentification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	post "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(address string) user.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user.NewUserServiceClient(conn)
}

func NewAuthentificationClient(address string) authentification.AuthentificationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return authentification.NewAuthentificationServiceClient(conn)
}

func NewPostClient(address string) post.PostServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}
	return post.NewPostServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

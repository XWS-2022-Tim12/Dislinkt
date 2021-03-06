package services

import (
	"log"

	authentification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	notification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/notification_service"
	job "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	jobSuggestions "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	post "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	message "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/message_service"
	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	userSuggestions "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
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

func NewMessageClient(address string) message.MessageServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Message service: %v", err)
	}
	return message.NewMessageServiceClient(conn)
}

func NewJobClient(address string) job.JobServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Job service: %v", err)
	}
	return job.NewJobServiceClient(conn)
}

func NewJobSuggestionsClient(address string) jobSuggestions.JobSuggestionsServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}
	return jobSuggestions.NewJobSuggestionsServiceClient(conn)
}

func NewUserSuggestionsClient(address string) userSuggestions.UserSuggestionsServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User Suggestions service: %v", err)
	}
	return userSuggestions.NewUserSuggestionsServiceClient(conn)
}

func NewNotificationClient(address string) notification.NotificationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Notification service: %v", err)
	}
	return notification.NewNotificationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

package api

import (
	"context"
	"log"
	"os"

	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/tracer"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/application"
	otgo "github.com/opentracing/opentracing-go"
)

var (
    InfoLogger  *log.Logger
	ErrorLogger *log.Logger
    trace       otgo.Tracer
)

type UserSuggestionsHandler struct {
	pb.UnimplementedUserSuggestionsServiceServer
	service *application.UserSuggestionsService
}

func NewUserSuggestionsHandler(service *application.UserSuggestionsService) *UserSuggestionsHandler {
	return &UserSuggestionsHandler{
		service: service,
	}
}

func init() {
    trace, _ = tracer.Init("user-suggestions-service")
    otgo.SetGlobalTracer(trace)
    infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

    errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
    if err1 != nil {
        log.Fatal(err1)
    }
    ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *UserSuggestionsHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapNewUser(request.User)
	_, err := handler.service.Register(user)
	if err != nil {
		ErrorLogger.Println("Action: 43, Message: Can not register user suggestion!")
		return nil, err
	}
	InfoLogger.Println("Action: 44, Message: User suggestion registered successfully!")
	return &pb.RegisterResponse{}, nil
}

func (handler *UserSuggestionsHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	Suggestions, err := handler.service.GetAll()
	if err != nil {
		ErrorLogger.Println("Action: 45, Message: Can not retrieve user suggestions!")
		return nil, err
	}
	InfoLogger.Println("Action: 46, Message: User suggestions retrieved successfully!")

	response := &pb.GetAllResponse{}

	for _, suggestion := range Suggestions {
		current := mapUser(suggestion)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

package api

import (
	"context"
	"log"
	"os"

	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/tracer"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/application"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	otgo "github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
    InfoLogger  *log.Logger
	ErrorLogger *log.Logger
    trace       otgo.Tracer
)

type SessionHandler struct {
	pb.UnimplementedAuthentificationServiceServer
	service *application.SessionService
}

func NewSessionHandler(service *application.SessionService) *SessionHandler {
	sh := &SessionHandler{
		service: service,
	}
	return sh
}

func init() {
    trace, _ = tracer.Init("authentification-service")
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

func (handler *SessionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	session, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	sessionPb := mapSession(session)
	response := &pb.GetResponse{
		Session: sessionPb,
	}
	return response, nil
}

func (handler *SessionHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	sessions, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Sessions: []*pb.Session{},
	}
	for _, session := range sessions {
		current := mapSession(session)
		response.Sessions = append(response.Sessions, current)
	}
	return response, nil
}

func (handler *SessionHandler) GetByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {
	id := request.UserId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	session, err := handler.service.GetByUserId(objectId)
	if err != nil {
		return nil, err
	}
	sessionPb := mapSession(session)
	response := &pb.GetByUserIdResponse{
		Session: sessionPb,
	}
	return response, nil
}

func (handler *SessionHandler) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	session := mapNewSession(request.Session)
	successs, err := handler.service.Add(session)
	if err != nil {
		ErrorLogger.Println("Action: 1, Message: Can not add session!")
        return nil, err
	}
	InfoLogger.Println("Action: 2, Message: Session added successfully!")
	response := &pb.AddResponse{
		Success: successs,
	}
	return response, err
}

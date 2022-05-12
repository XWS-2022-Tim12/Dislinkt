package api

import (
	"context"

	"github.com/XWS-2022-Tim12/Dislinkt/agentska/authentification_service/application"
	pb "github.com/XWS-2022-Tim12/Dislinkt/agentska/common/proto/authentification_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionHandler struct {
	pb.UnimplementedAuthentificationServiceServer
	service *application.SessionService
}

func NewSessionHandler(service *application.SessionService) *SessionHandler {
	return &SessionHandler{
		service: service,
	}
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

func (handler *SessionHandler) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	session := mapNewSession(request.Session)
	successs, err := handler.service.Add(session)
	response := &pb.AddResponse{
		Success: successs,
	}
	return response, err
}

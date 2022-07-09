package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/application"
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

func (handler *UserSuggestionsHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapNewUser(request.User)
	_, err := handler.service.Register(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (handler *UserSuggestionsHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	Suggestions, _ := handler.service.GetAll()

	response := &pb.GetAllResponse{}

	for _, suggestion := range Suggestions {
		current := mapUser(suggestion)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

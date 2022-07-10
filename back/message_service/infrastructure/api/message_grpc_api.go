package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/message_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer
	service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (handler *MessageHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	message, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	messagePb := mapMessage(message)
	response := &pb.GetResponse{
		Message: messagePb,
	}
	return response, nil
}

func (handler *MessageHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	messages, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messages {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

func (handler *MessageHandler) AddNewMessage(ctx context.Context, request *pb.AddNewMessageRequest) (*pb.AddNewMessageResponse, error) {
	message := mapNewMessage(request.Message)
	successs, err := handler.service.AddNewMessage(message)
	response := &pb.AddNewMessageResponse{
		Success: successs,
	}
	return response, err
}

func (handler *MessageHandler) GetMessagesBySenderAndReceiver(ctx context.Context, request *pb.GetMessagesBySenderAndReceiverRequest) (*pb.GetMessagesBySenderAndReceiverResponse, error) {
	sender := request.Sender
	receiver := request.Receiver
	messagesBySenderAndReceiver, err := handler.service.GetMessagesBySenderAndReceiver(sender, receiver)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMessagesBySenderAndReceiverResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messagesBySenderAndReceiver {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

func (handler *MessageHandler) GetMessagesByUsername(ctx context.Context, request *pb.GetMessagesByUsernameRequest) (*pb.GetMessagesByUsernameResponse, error) {
	username := request.Username
	messagesByUsername, err := handler.service.GetMessagesByUsername(username)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMessagesByUsernameResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messagesByUsername {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}
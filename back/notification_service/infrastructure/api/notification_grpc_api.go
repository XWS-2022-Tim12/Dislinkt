package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/notification_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	service *application.NotificationService
}

func NewNotificationHandler(service *application.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (handler *NotificationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	notification, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	notificationPb := mapNotification(notification)
	response := &pb.GetResponse{
		Notification: notificationPb,
	}
	return response, nil
}

func (handler *NotificationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	notifications, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) SearchBySender(ctx context.Context, request *pb.SearchBySenderRequest) (*pb.SearchBySenderResponse, error) {
	content := request.Sender
	notifications, err := handler.service.SearchBySender(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchBySenderResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) SearchByReceiver(ctx context.Context, request *pb.SearchByReceiverRequest) (*pb.SearchByReceiverResponse, error) {
	content := request.Receiver
	notifications, err := handler.service.SearchByReceiver(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByReceiverResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) SearchByNotificationType(ctx context.Context, request *pb.SearchByNotificationTypeRequest) (*pb.SearchByNotificationTypeResponse, error) {
	content := request.NotificationType
	notifications, err := handler.service.SearchByNotificationType(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByNotificationTypeResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	notification := mapNewNotification(request.Notification)
	successs, err := handler.service.Add(notification)
	response := &pb.AddResponse{
		Success: successs,
	}
	return response, err
}

func (handler *NotificationHandler) Edit(ctx context.Context, request *pb.EditRequest) (*pb.EditResponse, error) {
	notification := mapChangesOfNotification(request.Notification)
	successs, err := handler.service.Edit(notification)
	response := &pb.EditResponse{
		Success: successs,
	}
	return response, err
}

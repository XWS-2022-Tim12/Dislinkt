package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/notification_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapNotification(notification *domain.Notification) *pb.Notification {
	notificationPb := &pb.Notification{
		Id:                 notification.Id.Hex(),
		Sender:        		notification.Sender,
		Receiver:           notification.Receiver,
		CreationDate:       timestamppb.New(notification.CreationDate),
		NotificationType:   notification.NotificationType,
		Description:        notification.Description,
		IsRead:           	notification.IsRead,
	}
	return notificationPb
}

func mapNewNotification(notificationPb *pb.Notification) *domain.Notification {
	notification := &domain.Notification{
		Id:                 primitive.NewObjectID(),
		Sender:        		notificationPb.Sender,
		Receiver:           notificationPb.Receiver,
		CreationDate:       time.Now(),
		NotificationType:   notificationPb.NotificationType,
		Description:        notificationPb.Description,
		IsRead:           	notificationPb.IsRead,
	}
	return notification
}

func mapChangesOfNotification(notificationPb *pb.Notification) *domain.Notification {
	id, _ := primitive.ObjectIDFromHex(notificationPb.Id)

	notification := &domain.Notification{
		Id:                 id,
		Sender:        		notificationPb.Sender,
		Receiver:           notificationPb.Receiver,
		CreationDate:       notificationPb.CreationDate.AsTime(),
		NotificationType:   notificationPb.NotificationType,
		Description:        notificationPb.Description,
		IsRead:           	notificationPb.IsRead,
	}
	return notification
}

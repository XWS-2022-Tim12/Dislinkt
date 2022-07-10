package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/message_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapMessage(message *domain.Message) *pb.Message {
	messagePb := &pb.Message{
		Id:       message.Id.Hex(),
		Text:     message.Text,
		Date:	  timestamppb.New(message.Date),
		SenderUsername: message.SenderUsername,
		ReceiverUsername: message.ReceiverUsername,
	}
	return messagePb
}

func mapNewMessage(messagePb *pb.Message) *domain.Message {
	if messagePb.Date != nil {
		message := &domain.Message{
			Id:       primitive.NewObjectID(),
			Text:     messagePb.Text,
			Date: 	  messagePb.Date.AsTime(),
			SenderUsername: messagePb.SenderUsername,
			ReceiverUsername: messagePb.ReceiverUsername,
		}
		return message
	} else {
		message := &domain.Message{
			Id:       primitive.NewObjectID(),
			Text:     messagePb.Text,
			Date: 	  time.Now(),
			SenderUsername: messagePb.SenderUsername,
			ReceiverUsername: messagePb.ReceiverUsername,
		}
		return message
	}
}
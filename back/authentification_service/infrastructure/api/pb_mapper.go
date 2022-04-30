package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapSession(session *domain.Session) *pb.Session {
	sessionPb := &pb.Session{
		Id:           session.Id.Hex(),
		UserId:       session.UserId.Hex(),
		Date:         timestamppb.New(session.Date),
	}
	return sessionPb
}

func mapNewSession(sessionPb *pb.Session) *domain.Session {
	session := &domain.Session{
		Id:           primitive.NewObjectID(),
		UserId        primitive.NewObjectID(),
		Date:     time.Now(),
	}
	return session
}
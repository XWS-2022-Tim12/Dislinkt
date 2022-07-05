package api

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	events "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/register_user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapSession(session *domain.Session) *pb.Session {
	sessionPb := &pb.Session{
		Id:     session.Id.Hex(),
		UserId: session.UserId.Hex(),
		Date:   timestamppb.New(session.Date),
		Role:   session.Role,
	}
	return sessionPb
}

func mapNewSession(sessionPb *pb.Session) *domain.Session {
	userId, _ := primitive.ObjectIDFromHex(sessionPb.UserId)

	session := &domain.Session{
		Id:     primitive.NewObjectID(),
		UserId: userId,
		Date:   time.Now(),
		Role:   sessionPb.Role,
	}
	return session
}

func mapCommandToSession(regCommand *events.RegisterUserCommand) *domain.Session {
	userId, _ := primitive.ObjectIDFromHex(regCommand.User.Id)
	session := &domain.Session{
		Id:     primitive.NewObjectID(),
		UserId: userId,
		Date:   time.Now(),
		Role:   "user",
	}
	return session
}

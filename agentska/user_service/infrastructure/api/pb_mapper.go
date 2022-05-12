package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/agentska/common/proto/user_service"
	"github.com/XWS-2022-Tim12/Dislinkt/agentska/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:           user.Id.Hex(),
		Firstname:    user.Firstname,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Gender:       mapGender(user.Gender),
		BirthDay:     timestamppb.New(user.BirthDay),
		Username:     user.Username,
		Password:     user.Password,
	}
	return userPb
}

func mapNewUser(userPb *pb.User) *domain.User {

	if userPb.BirthDay != nil {
		user := &domain.User{
			Id:           primitive.NewObjectID(),
			Firstname:    userPb.Firstname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     userPb.BirthDay.AsTime(),
			Username:     userPb.Username,
			Password:     userPb.Password,
		}
		return user
	} else {
		user := &domain.User{
			Id:           primitive.NewObjectID(),
			Firstname:    userPb.Firstname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     time.Now(),
			Username:     userPb.Username,
			Password:     userPb.Password,
		}
		return user
	}
}

func mapGender(status domain.GenderEnum) pb.User_GenderEnum {
	switch status {
	case domain.Male:
		return pb.User_Male
	}
	return pb.User_Female

}

func mapNewGender(status pb.User_GenderEnum) domain.GenderEnum {
	switch status {
	case pb.User_Male:
		return domain.Male
	}
	return domain.Female

}

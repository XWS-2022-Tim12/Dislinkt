package api

import (
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
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
		Biography:    user.Biography,
		Experience:   user.Experience,
		Education:    mapEducation(user.Education),
		Skills:       user.Skills,
		Interests:    user.Interests,
		Password:     user.Password,
	}
	return userPb
}

func mapNewUser(userPb *pb.User) *domain.User {

	user := &domain.User{
		Id:           primitive.NewObjectID(),
		Firstname:    userPb.Firstname,
		Email:        userPb.Email,
		MobileNumber: userPb.MobileNumber,
		Gender:       mapNewGender(userPb.Gender),
		BirthDay:     userPb.BirthDay.AsTime(),
		Username:     userPb.Username,
		Biography:    userPb.Biography,
		Experience:   userPb.Experience,
		Education:    mapNewEducation(userPb.Education),
		Skills:       userPb.Skills,
		Interests:    userPb.Interests,
		Password:     userPb.Password,
	}

	return user
}

func mapEducation(status domain.EducationEnum) pb.User_EducationEnum {
	switch status {
	case domain.PrimaryEducation:
		return pb.User_PrimaryEducation
	case domain.LowerSecondaryEducation:
		return pb.User_LowerSecondaryEducation
	case domain.UpperSecondaryEducation:
		return pb.User_UpperSecondaryEducation
	case domain.Bachelor:
		return pb.User_Bachelor
	case domain.Master:
		return pb.User_Master
	}
	return pb.User_Doctorate

}

func mapNewEducation(status pb.User_EducationEnum) domain.EducationEnum {
	switch status {
	case pb.User_PrimaryEducation:
		return domain.PrimaryEducation
	case pb.User_LowerSecondaryEducation:
		return domain.LowerSecondaryEducation
	case pb.User_UpperSecondaryEducation:
		return domain.UpperSecondaryEducation
	case pb.User_Bachelor:
		return domain.Bachelor
	case pb.User_Master:
		return domain.Master
	}
	return domain.Doctorate

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

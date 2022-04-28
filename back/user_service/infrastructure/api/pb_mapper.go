package api

import (
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
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

func mapGender(status domain.GenderEnum) pb.User_GenderEnum {
	switch status {
	case domain.Male:
		return pb.User_Male
	}
	return pb.User_Female

}

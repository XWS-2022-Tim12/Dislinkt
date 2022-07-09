package api

import (
	"strconv"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain"
)

func mapUser(userPb *domain.User) *pb.User {
	x := int(userPb.Id)
	user := &pb.User{
		Id:        strconv.Itoa(x),
		FirstName: userPb.FirstName,
		Email:     userPb.Email,
		Username:  userPb.Username,
		Interests: userPb.Interests,
	}
	return user
}

func mapNewUser(userPb *pb.User) *domain.User {
	user := &domain.User{
		FirstName: userPb.FirstName,
		Email:     userPb.Email,
		Username:  userPb.Username,
		Interests: userPb.Interests,
	}
	return user
}

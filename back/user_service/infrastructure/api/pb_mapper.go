package api

import (
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       user.Id.Hex(),
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	return userPb
}

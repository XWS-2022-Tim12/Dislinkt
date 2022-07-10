package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	uh := &UserHandler{
		service: service,
	}
	return uh
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetByUsername(ctx context.Context, request *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	username := request.Username

	user, err := handler.service.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetByUsernameResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetPublicUserByUsername(ctx context.Context, request *pb.GetPublicUserByUsernameRequest) (*pb.GetPublicUserByUsernameResponse, error) {
	username := request.Username

	user, err := handler.service.GetPublicUserByUsername(username)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetPublicUserByUsernameResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) GetAllPublicUsers(ctx context.Context, request *pb.GetAllPublicUsersRequest) (*pb.GetAllPublicUsersResponse, error) {
	publicUsers, err := handler.service.GetAllPublicUsers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllPublicUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range publicUsers {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) GetAllPublicUsersByUsername(ctx context.Context, request *pb.GetAllPublicUsersByUsernameRequest) (*pb.GetAllPublicUsersByUsernameResponse, error) {
	username := request.Username
	publicUsersByUsername, err := handler.service.GetAllPublicUsersByUsername(username)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllPublicUsersByUsernameResponse{
		Users: []*pb.User{},
	}
	for _, user := range publicUsersByUsername {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) GetFollowingNotBlockedUsers(ctx context.Context, request *pb.GetFollowingNotBlockedUsersRequest) (*pb.GetFollowingNotBlockedUsersResponse, error) {
	username := request.Username
	followingNotBlockedUsers, err := handler.service.GetFollowingNotBlockedUsers(username)
	if err != nil {
		return nil, err
	}

	response := &pb.GetFollowingNotBlockedUsersResponse{
		Users: []*pb.User{},
	}
	for _, user := range followingNotBlockedUsers {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapNewUser(request.User)
	successs, err := handler.service.Register(user)
	response := &pb.RegisterResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) UpdateBasicInfo(ctx context.Context, request *pb.UpdateBasicInfoRequest) (*pb.UpdateBasicInfoResponse, error) {
	user := mapBasicInfoUser(request.User)
	successs, err := handler.service.UpdateBasicInfo(user)
	response := &pb.UpdateBasicInfoResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) UpdateAdvancedInfo(ctx context.Context, request *pb.UpdateAdvancedInfoRequest) (*pb.UpdateAdvancedInfoResponse, error) {
	user := mapAdvancedInfoUser(request.User)
	successs, err := handler.service.UpdateAdvancedInfo(user)
	response := &pb.UpdateAdvancedInfoResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) UpdatePersonalInfo(ctx context.Context, request *pb.UpdatePersonalInfoRequest) (*pb.UpdatePersonalInfoResponse, error) {
	user := mapPersonalInfoUser(request.User)
	successs, err := handler.service.UpdatePersonalInfo(user)
	response := &pb.UpdatePersonalInfoResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) UpdateAllInfo(ctx context.Context, request *pb.UpdateAllInfoRequest) (*pb.UpdateAllInfoResponse, error) {
	user := mapAllInfoUser(request.User)
	successs, err := handler.service.UpdateAllInfo(user)
	response := &pb.UpdateAllInfoResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) FollowPublicProfile(ctx context.Context, request *pb.FollowPublicProfileRequest) (*pb.FollowPublicProfileResponse, error) {
	user := mapUserToFollow(request.User)
	successs, err := handler.service.FollowPublicProfile(user)
	response := &pb.FollowPublicProfileResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) AcceptFollowingRequest(ctx context.Context, request *pb.AcceptFollowingRequestRequest) (*pb.AcceptFollowingRequestResponse, error) {
	user := mapUserToFollow(request.User)
	successs, err := handler.service.AcceptFollowingRequest(user)
	response := &pb.AcceptFollowingRequestResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) RejectFollowingRequest(ctx context.Context, request *pb.RejectFollowingRequestRequest) (*pb.RejectFollowingRequestResponse, error) {
	user := mapUserToFollow(request.User)
	successs, err := handler.service.RejectFollowingRequest(user)
	response := &pb.RejectFollowingRequestResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) BlockUser(ctx context.Context, request *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {
	user := mapUserToBlock(request.User)
	successs, err := handler.service.BlockUser(user)
	response := &pb.BlockUserResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) ChangeNotifications(ctx context.Context, request *pb.ChangeNotificationsRequest) (*pb.ChangeNotificationsResponse, error) {
	user := mapUserToChangeNotifications(request.User)
	successs, err := handler.service.ChangeNotifications(user)
	response := &pb.ChangeNotificationsResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) ChangeNotificationsUsers(ctx context.Context, request *pb.ChangeNotificationsUsersRequest) (*pb.ChangeNotificationsUsersResponse, error) {
	user := mapUserToChangeNotifications(request.User)
	successs, err := handler.service.ChangeNotificationsUsers(user)
	response := &pb.ChangeNotificationsUsersResponse{
		Success: successs,
	}
	return response, err
}

func (handler *UserHandler) ChangeNotificationsMessages(ctx context.Context, request *pb.ChangeNotificationsMessagesRequest) (*pb.ChangeNotificationsMessagesResponse, error) {
	user := mapUserToChangeNotifications(request.User)
	successs, err := handler.service.ChangeNotificationsMessages(user)
	response := &pb.ChangeNotificationsMessagesResponse{
		Success: successs,
	}
	return response, err
}

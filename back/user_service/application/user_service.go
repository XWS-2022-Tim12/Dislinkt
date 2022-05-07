package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Register(user *domain.User) (string, error) {
	success, err := service.store.Insert(user)
	return success, err
}

func (service *UserService) UpdateBasicInfo(user *domain.User) (string, error) {
	success, err := service.store.UpdateBasicInfo(user)
	return success, err
}

func (service *UserService) UpdateAdvancedInfo(user *domain.User) (string, error) {
	success, err := service.store.UpdateAdvancedInfo(user)
	return success, err
}

func (service *UserService) UpdatePersonalInfo(user *domain.User) (string, error) {
	success, err := service.store.UpdatePersonalInfo(user)
	return success, err
}

func (service *UserService) UpdateAllInfo(user *domain.User) (string, error) {
	success, err := service.store.UpdateAllInfo(user)
	return success, err
}

func (service *UserService) FollowPublicProfile(user *domain.User) (string, error) {
	success, err := service.store.FollowPublicProfile(user)
	return success, err
}

func (service *UserService) AcceptFollowingRequest(user *domain.User) (string, error) {
	success, err := service.store.AcceptFollowingRequest(user)
	return success, err
}
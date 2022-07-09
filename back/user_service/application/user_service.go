package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store        domain.UserStore
	orchestrator *AddUserOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *AddUserOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}

func (service *UserService) GetPublicUserByUsername(username string) (*domain.User, error) {
	return service.store.GetPublicUserByUsername(username)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) GetAllPublicUsers() ([]*domain.User, error) {
	return service.store.GetAllPublicUsers()
}

func (service *UserService) GetAllPublicUsersByUsername(username string) ([]*domain.User, error) {
	return service.store.GetAllPublicUsersByUsername(username)
}

func (service *UserService) Register(user *domain.User) (string, error) {
	success, err := service.store.Insert(user)
	if err != nil {
		return success, err
	}
	err = service.orchestrator.Start(user)
	if err != nil {
		return success, nil
	}

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

func (service *UserService) RejectFollowingRequest(user *domain.User) (string, error) {
	success, err := service.store.RejectFollowingRequest(user)
	return success, err
}

func (service *UserService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *UserService) BlockUser(user *domain.User) (string, error) {
	success, err := service.store.BlockUser(user)
	return success, err
}

func (service *UserService) ChangeNotifications(user *domain.User) (string, error) {
	success, err := service.store.ChangeNotifications(user)
	return success, err
}

func (service *UserService) ChangeNotificationsUsers(user *domain.User) (string, error) {
	success, err := service.store.ChangeNotificationsUsers(user)
	return success, err
}

func (service *UserService) ChangeNotificationsMessages(user *domain.User) (string, error) {
	success, err := service.store.ChangeNotificationsMessages(user)
	return success, err
}

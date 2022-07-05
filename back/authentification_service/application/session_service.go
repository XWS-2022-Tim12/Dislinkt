package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionService struct {
	store domain.SessionStore
}

func NewSessionService(store domain.SessionStore) *SessionService {
	return &SessionService{
		store: store,
	}
}

func (service *SessionService) Get(id primitive.ObjectID) (*domain.Session, error) {
	return service.store.Get(id)
}
func (service *SessionService) GetAll() ([]*domain.Session, error) {
	return service.store.GetAll()
}
func (service *SessionService) GetByUserId(id primitive.ObjectID) (*domain.Session, error) {
	return service.store.GetByUserId(id)
}

func (service *SessionService) Add(session *domain.Session) (string, error) {
	success, err := service.store.Insert(session)
	return success, err
}

func (service *SessionService) DeleteByUserId(userId string) error {
	err := service.store.DeleteByUserId(userId)
	return err
}

package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store domain.MessageStore
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) Get(id primitive.ObjectID) (*domain.Message, error) {
	return service.store.Get(id)
}

func (service *MessageService) GetAll() ([]*domain.Message, error) {
	return service.store.GetAll()
}

func (service *MessageService) AddNewMessage(message *domain.Message) (string, error) {
	success, err := service.store.Insert(message)
	return success, err
}

func (service *MessageService) GetMessagesBySenderAndReceiver(user1, user2 string) ([]*domain.Message, error) {
	return service.store.GetMessagesBySenderAndReceiver(user1, user2)
}

func (service *MessageService) GetMessagesByUsername(username string) ([]*domain.Message, error) {
	return service.store.GetMessagesByUsername(username)
}
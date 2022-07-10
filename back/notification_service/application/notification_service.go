package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationService struct {
	store domain.NotificationStore
}

func NewNotificationService(store domain.NotificationStore) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (service *NotificationService) Get(id primitive.ObjectID) (*domain.Notification, error) {
	return service.store.Get(id)
}

func (service *NotificationService) GetAll() ([]*domain.Notification, error) {
	return service.store.GetAll()
}

func (service *NotificationService) SearchBySender(content string) ([]*domain.Notification, error) {
	return service.store.SearchBySender(content)
}

func (service *NotificationService) SearchByReceiver(content string) ([]*domain.Notification, error) {
	return service.store.SearchByReceiver(content)
}

func (service *NotificationService) SearchByNotificationType(content string) ([]*domain.Notification, error) {
	return service.store.SearchByNotificationType(content)
}

func (service *NotificationService) Add(notification *domain.Notification) (string, error) {
	success, err := service.store.Insert(notification)
	return success, err
}

func (service *NotificationService) Edit(notification *domain.Notification) (string, error) {
	success, err := service.store.Edit(notification)
	return success, err
}

package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationStore interface {
	Get(id primitive.ObjectID) (*Notification, error)
	GetAll() ([]*Notification, error)
	SearchBySender(content string) ([]*Notification, error)
	SearchByReceiver(content string) ([]*Notification, error)
	SearchByNotificationType(content string) ([]*Notification, error)
	Insert(notification *Notification) (string, error)
	Edit(notification *Notification) (string, error)
	DeleteAll()
}

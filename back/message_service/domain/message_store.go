package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStore interface {
	Get(id primitive.ObjectID) (*Message, error)
	GetAll() ([]*Message, error)
	GetMessagesBySenderAndReceiver(string, string) ([]*Message, error)
	Insert(message *Message) (string, error)
	DeleteAll()
	GetMessagesByUsername(string) ([]*Message, error)
}

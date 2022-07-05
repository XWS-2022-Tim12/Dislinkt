package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionStore interface {
	Get(id primitive.ObjectID) (*Session, error)
	GetAll() ([]*Session, error)
	GetByUserId(id primitive.ObjectID) (*Session, error)
	Insert(session *Session) (string, error)
	DeleteAll()
	Delete(id primitive.ObjectID)
	DeleteByUserId(userId string) error
}

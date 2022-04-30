package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionStore interface {
	Get(id primitive.ObjectID) (*Session, error)
	Insert(session *Session) (string, error)
	DeleteAll()
}

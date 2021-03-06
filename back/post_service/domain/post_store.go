package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	GetUserPosts(username string) ([]*Post, error)
	Insert(post *Post) (string, error)
	UpdatePost(post *Post) (string, error)
	DeleteAll()
}

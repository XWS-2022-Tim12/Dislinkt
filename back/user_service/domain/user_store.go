package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAllPublicUsers() ([]*User, error)
	GetAllUsersByUsername(username string) ([]*User, error)
	GetAllPublicUsersByUsername(username string) ([]*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) (string, error)
	UpdateBasicInfo(user *User) (string, error)
	UpdateAdvancedInfo(user *User) (string, error)
	UpdatePersonalInfo(user *User) (string, error)
	UpdateAllInfo(user *User) (string, error)
	DeleteAll()
	FollowPublicProfile(user *User) (string, error)
}

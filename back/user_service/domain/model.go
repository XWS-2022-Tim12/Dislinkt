package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

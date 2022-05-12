package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"userId"`
	Date        time.Time          `bson:"date"`
	Role        string             `bson:"role"`
}
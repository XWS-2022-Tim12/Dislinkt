package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       primitive.ObjectID `bson:"userId"`
	CreationDay  time.Time          `bson:"creationDay"`
	Position     string             `bson:"position"`
	Description  string             `bson:"description"`
	Requirements string             `bson:"requirements"`
}

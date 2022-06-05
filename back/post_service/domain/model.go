package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Date     time.Time          `bson:"date"`
	Likes    int32              `bson:"likes"`
	Dislikes int32              `bson:"dislikes"`
	Comments []string           `bson:"comments"`
	Username string             `bson:"username"`
	ImageContent string			`bson:"imageContent"`
}

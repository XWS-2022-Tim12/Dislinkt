package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Likes    int32              `bson:"likes"`
	Dislikes int32              `bson:"dislikes"`
	Comments []string           `bson:"comments"`
	Username string             `bson:"username"`
	ImageContent string			`bson:"imageContent"`
}

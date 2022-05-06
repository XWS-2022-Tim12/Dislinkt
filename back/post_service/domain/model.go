package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EducationEnum int8

type GenderEnum int8

const (
	PrimaryEducation EducationEnum = iota
	LowerSecondaryEducation
	UpperSecondaryEducation
	Bachelor
	Master
	Doctorate
)

type Post struct {
	Id    primitive.ObjectID `bson:"_id"`
	Text  string             `bson:"text"`
	Image string             `bson:"image"`
	Link  string             `bson:"link"`
}

package domain

import (
	"time"

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

const (
	Male GenderEnum = iota
	Female
)

func (status EducationEnum) String() string {
	switch status {
	case PrimaryEducation:
		return "Primary education"
	case LowerSecondaryEducation:
		return "Lower secondary education"
	case UpperSecondaryEducation:
		return "Upper secondary education"
	case Bachelor:
		return "Bachelor"
	case Master:
		return "Master"
	case Doctorate:
		return "Doctorate"
	}
	return "Unknown"
}

func (status GenderEnum) String() string {
	switch status {
	case Male:
		return "Male"
	case Female:
		return "Female"
	}
	return "Unknown"
}

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Firstname    string             `bson:"firstname"`
	Email        string             `bson:"email"`
	MobileNumber string             `bson:"mobileNumber"`
	Gender       GenderEnum         `bson:"gender"`
	BirthDay     time.Time          `bson:"birthDay"`
	Username     string             `bson:"username"`
	Biography    string             `bson:"biography"`
	Experience   string             `bson:"experience"`
	Education    EducationEnum      `bson:"education"`
	Skills       string             `bson:"skills"`
	Interests    string             `bson:"interests"`
	Password     string             `bson:"password"`
}
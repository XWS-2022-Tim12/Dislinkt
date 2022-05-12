package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/agentska/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:           getObjectId("62fghcc3a34d25d8567f9f82"),
		Firstname:    "Prvi",
		Email:        "prvi@gmail.com",
		MobileNumber: "05654127",
		Gender:       domain.Male,
		BirthDay:     time.Now(),
		Username:     "mico",
		Password:     "prvasifra",
	},
	{
		Id:           getObjectId("62fsfag3a34d25d8567f9f83"),
		Firstname:    "Drugi",
		Email:        "drugi@gmail.com",
		MobileNumber: "067415402",
		Gender:       domain.Female,
		BirthDay:     time.Now(),
		Username:     "nina",
		Password:     "drugasifra",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:           		getObjectId("623b0cc3a34d25d8567f9f82"),
		Firstname:    		"ime",
		Email:        		"email@gmail.com",
		MobileNumber: 		"02302303",
		Gender:       		domain.Male,
		BirthDay:     		time.Now(),
		Username:     		"mico",
		Biography:    		"Zivio u Gacku",
		Experience:   		"Radio na farmi",
		Education:    		domain.Master,
		Skills:       		"Programiranje",
		Interests:    		"Programiranje",
		Password:     		"nekasifra",
		FollowingUsers: 	[]string{
			"nina",
		},
		FollowedByUsers:	[]string{},
		Public:       		true,
	},
	{
		Id:             	getObjectId("623b0cc3a34d25d8567f9f83"),
		Firstname:      	"ime2",
		Email:          	"imejl@gmail.com",
		MobileNumber:   	"932939332",
		Gender:         	domain.Female,
		BirthDay:       	time.Now(),
		Username:       	"nina",
		Biography:      	"Neka biografija",
		Experience:     	"Radila na motorima",
		Education:      	domain.Bachelor,
		Skills:         	"Mehanika",
		Interests:      	"Mehanika",
		Password:       	"nekasifra2",
		FollowingUsers: 	[]string{},
		FollowedByUsers: 	[]string{
			"mico",
		},
		Public:       		true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

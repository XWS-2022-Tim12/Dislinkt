package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
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
		Biography:    "Zivio u Gacku",
		Experience:   "Radio na farmi",
		Education:    domain.Master,
		Skills:       "Programiranje",
		Interests:    "Programiranje",
		Password:     "prvasifra",
		FollowingUsers: []string{
			"nina",
			"treci",
		},
		FollowedByUsers:   []string{},
		FollowingRequests: []string{},
		Public:            true,
	},
	{
		Id:             getObjectId("62fsfag3a34d25d8567f9f83"),
		Firstname:      "Drugi",
		Email:          "drugi@gmail.com",
		MobileNumber:   "067415402",
		Gender:         domain.Female,
		BirthDay:       time.Now(),
		Username:       "nina",
		Biography:      "Neka biografija",
		Experience:     "Radila na motorima",
		Education:      domain.Master,
		Skills:         "Programiranje",
		Interests:      "Programiranje",
		Password:       "drugasifra",
		FollowingUsers: []string{},
		FollowedByUsers: []string{
			"mico",
		},
		FollowingRequests: []string{},
		Public:            true,
	},
	{
		Id:             getObjectId("62Ssafc3a34d25d8567f9f84"),
		Firstname:      "Treci",
		Email:          "treci@gmail.com",
		MobileNumber:   "062541758",
		Gender:         domain.Female,
		BirthDay:       time.Now(),
		Username:       "treci",
		Biography:      "Neka biografija",
		Experience:     "Radila svasta",
		Education:      domain.Bachelor,
		Skills:         "Skijaska licenca",
		Interests:      "Skijanje",
		Password:       "trecasifra",
		FollowingUsers: []string{},
		FollowedByUsers: []string{
			"mico",
		},
		FollowingRequests: []string{},
		Public:            false,
	},
	{
		Id:                getObjectId("6fds7c3a34d251f567f9f85"),
		Firstname:         "Cetvrti",
		Email:             "cetvrti@gmail.com",
		MobileNumber:      "064257136",
		Gender:            domain.Male,
		BirthDay:          time.Now(),
		Username:          "cetvrti",
		Biography:         "Zivio u Beogradu",
		Experience:        "Radio pri vojsci",
		Education:         domain.Master,
		Skills:            "Zna sa puskom",
		Interests:         "Rat",
		Password:          "cetvrtasifra",
		FollowingUsers:    []string{},
		FollowedByUsers:   []string{},
		FollowingRequests: []string{},
		Public:            false,
	},
	{
		Id:                getObjectId("fdf34d25d8567f9f86"),
		Firstname:         "Peti",
		Email:             "peti@gmail.com",
		MobileNumber:      "057485264",
		Gender:            domain.Male,
		BirthDay:          time.Now(),
		Username:          "peti",
		Biography:         "Zivio u Novom Sadu",
		Experience:        "Profesorsko iskustvo",
		Education:         domain.Master,
		Skills:            "Predavanje",
		Interests:         "Skola",
		Password:          "petasifra",
		FollowingUsers:    []string{},
		FollowedByUsers:   []string{},
		FollowingRequests: []string{},
		Public:            false,
	},
	{
		Id:                getObjectId("4ds834d25d8567f9f87"),
		Firstname:         "Sesti",
		Email:             "sesti@gmail.com",
		MobileNumber:      "063485264",
		Gender:            domain.Male,
		BirthDay:          time.Now(),
		Username:          "sesti",
		Biography:         "Zivio u Nisu",
		Experience:        "Sportsko iskustvo",
		Education:         domain.Master,
		Skills:            "Trener",
		Interests:         "Sport",
		Password:          "sestasifra",
		FollowingUsers:    []string{},
		FollowedByUsers:   []string{},
		FollowingRequests: []string{},
		Public:            false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

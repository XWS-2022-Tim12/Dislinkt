package startup

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("523b0cc3a34d25d8567f9f82"),
		Text:     "Aaaaaaa",
		Image:    "",
		Link:     "",
		Likes:    14,
		Dislikes: 10,
		Comments: []string{
			"Lepa slika",
			"Top",
		},
		Username: "mico",
	},
	{
		Id:       getObjectId("524b0cc3a34d25d8567f9f82"),
		Text:     "Bbbbbbb",
		Image:    "",
		Link:     "",
		Likes:    44,
		Dislikes: 3,
		Comments: []string{
			"Sjajno",
		},
		Username: "nina",
	},
	{
		Id:       getObjectId("524bfksafk4d25d8567f9f82"),
		Text:     "New here",
		Image:    "",
		Link:     "",
		Likes:    2,
		Dislikes: 0,
		Comments: []string{},
		Username: "treci",
	},
	{
		Id:       getObjectId("v43cc3a34d25d8567f9f82"),
		Text:     "Neki tekst",
		Image:    "",
		Link:     "",
		Likes:    4,
		Dislikes: 3,
		Comments: []string{},
		Username: "cetvrti",
	},
	{
		Id:       getObjectId("t34v0cc3a34d25d8567f9f82"),
		Text:     "New post",
		Image:    "",
		Link:     "",
		Likes:    1,
		Dislikes: 1,
		Comments: []string{
			"Bravo",
			"Super",
		},
		Username: "treci",
	},
	{
		Id:       getObjectId("13410cc3a34d25d8567f9f82"),
		Text:     "Cao",
		Image:    "",
		Link:     "",
		Likes:    1,
		Dislikes: 10,
		Comments: []string{
			"Fuj",
			"Ne valja",
			"Glupo",
		},
		Username: "peti",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

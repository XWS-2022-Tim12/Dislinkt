package startup

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("523b0cc3a34d25d8567f9f82"),
		Text:     "Aaaaaaa",
		Likes:    14,
		Dislikes: 10,
		Comments: []string{
			"Lepa slika",
			"Top",
		},
		Username: "mico",
		ImageContent: "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",
	},
	{
		Id:       getObjectId("524b0cc3a34d25d8567f9f82"),
		Text:     "Bbbbbbb",
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
		Likes:    2,
		Dislikes: 0,
		Comments: []string{},
		Username: "treci",
	},
	{
		Id:       getObjectId("v43cc3a34d25d8567f9f82"),
		Text:     "Neki tekst",
		Likes:    4,
		Dislikes: 3,
		Comments: []string{},
		Username: "cetvrti",
	},
	{
		Id:       getObjectId("t34v0cc3a34d25d8567f9f82"),
		Text:     "New post",
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

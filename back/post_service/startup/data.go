package startup

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:    getObjectId("523b0cc3a34d25d8567f9f82"),
		Text:  "Aaaaaaa",
		Image: "",
		Link:  "",
	},
	{
		Id:    getObjectId("524b0cc3a34d25d8567f9f82"),
		Text:  "Bbbbbbb",
		Image: "",
		Link:  "",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

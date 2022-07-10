package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var messages = []*domain.Message{
	{
		Id:           			getObjectId("62caa2e95c5a04c8ce73f2b2"),
		Text:    	  			"Dobar dan",
		Date:           		time.Now(),
		SenderUsername: 		"mico",
		ReceiverUsername:       "nina",
	},
	{
		Id:           			getObjectId("62caa30b5c5a04c8ce73f2b4"),
		Text:    	  			"Halo",
		Date:           		time.Now(),
		SenderUsername: 		"mico",
		ReceiverUsername:       "treci",
	},
	{
		Id:           			getObjectId("62caa3535c5a04c8ce73f2b6"),
		Text:    	  			"halo",
		Date:           		time.Now(),
		SenderUsername: 		"nina",
		ReceiverUsername:       "mico",
	},
	{
		Id:           			getObjectId("62caa4925c5a04c8ce73f2b8"),
		Text:    	  			"zdravo",
		Date:           		time.Now(),
		SenderUsername: 		"treci",
		ReceiverUsername:       "nina",
	},
	{
		Id:           			getObjectId("62caa4ad5c5a04c8ce73f2ba"),
		Text:    	  			"zdravo",
		Date:           		time.Now(),
		SenderUsername: 		"treci",
		ReceiverUsername:       "mico",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
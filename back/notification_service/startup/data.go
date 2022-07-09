package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var notifications = []*domain.Notification {
	{
		Id:                 getObjectId("62fghcc3a34d25d8567f9f22"),
		Sender:           	"nina",
		Receiver:        	"mico",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"Description",
		IsRead:				true,
	},
	{
		Id:                 getObjectId("76f1yh1h378f134hfh34u8dh"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"Description",
		IsRead:				true,
	},
	{
		Id:                 getObjectId("75ef7eolj74hyt56tfd3er4v"),
		Sender:           	"nina",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"Description",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("12lkj8uhbg542sdc4rfv76y7"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "unfollow",
		Description:		"Description",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("46tgf54r45er4531s00b7884"),
		Sender:           	"peti",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "follow",
		Description:		"Description",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("378jh6tfrsx32vfro09h76tf"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "dislike",
		Description:		"Description",
		IsRead:				false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

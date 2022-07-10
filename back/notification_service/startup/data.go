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
		Description:		"User nina liked post from mico.",
		IsRead:				true,
	},
	{
		Id:                 getObjectId("76f1yh1h378f134hfh34u8dh"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"User treci liked post from cetvrti.",
		IsRead:				true,
	},
	{
		Id:                 getObjectId("75ef7eolj74hyt56tfd3er4v"),
		Sender:           	"nina",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"User nina liked post from cetvrti.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("12lkj8uhbg542sdc4rfv76y7"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "unfollow",
		Description:		"User treci unfollowed cetvrti.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("46tgf54r45er4531s00b7884"),
		Sender:           	"peti",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "follow",
		Description:		"User peti wants to follow cetvrti.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("378jh6tfrsx32vfro09h76tf"),
		Sender:           	"treci",
		Receiver:        	"cetvrti",
		CreationDate:       time.Now(),
		NotificationType:   "dislike",
		Description:		"User treci disliked post from cetvrti.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("578jh6tfrsx32vfro09h76tf"),
		Sender:           	"treci",
		Receiver:        	"nina",
		CreationDate:       time.Now(),
		NotificationType:   "like",
		Description:		"User treci liked post from nina.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("478jh6tfrsx32vfro09h76tf"),
		Sender:           	"cetvrti",
		Receiver:        	"treci",
		CreationDate:       time.Now(),
		NotificationType:   "dislike",
		Description:		"User cetvrti disliked post from treci.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("278jh6tfrsx32vfro09h76tf"),
		Sender:           	"sesti",
		Receiver:        	"peti",
		CreationDate:       time.Now(),
		NotificationType:   "acceptRequest",
		Description:		"User sesti acceptet following request from peti.",
		IsRead:				false,
	},
	{
		Id:                 getObjectId("178jh6tfrsx32vfro09h76tf"),
		Sender:           	"peti",
		Receiver:        	"mico",
		CreationDate:       time.Now(),
		NotificationType:   "comment",
		Description:		"User peti commented post from mico.",
		IsRead:				false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Sender             string 			  `bson:"sender"`
	Receiver           string             `bson:"receiver"`
	CreationDate       time.Time          `bson:"creationDate"`
	NotificationType   string             `bson:"notificationType"`
	Description        string             `bson:"description"`
	IsRead		       bool	              `bson:"isRead"`
}

package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	Date     time.Time          `bson:"date"`
	SenderUsername string       `bson:"senderUsername"`
	ReceiverUsername string		`bson:"receiverUsername"`
}

package persistence

import (
	"context"

	"github.com/XWS-2022-Tim12/Dislinkt/back/message_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "message"
	COLLECTION = "message"
)

type MessageMongoDBStore struct {
	messages *mongo.Collection
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	messages := client.Database(DATABASE).Collection(COLLECTION)
	return &MessageMongoDBStore{
		messages: messages,
	}
}

func (store *MessageMongoDBStore) Get(id primitive.ObjectID) (*domain.Message, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *MessageMongoDBStore) GetAll() ([]*domain.Message, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *MessageMongoDBStore) GetMessagesBySenderAndReceiver(user1, user2 string) ([]*domain.Message, error) {
	filter := bson.M{ "$or": bson.A{
						bson.M{"senderUsername": user1},
						bson.M{"senderUsername": user2},
					}, }
	messages, err := store.filter(filter)
	if err != nil {
		return nil, err
	}

	var messagesBySenderAndReceiver []*domain.Message
	for i := 0; i < len(messages); i++ {
		if (messages[i].SenderUsername == user1 && messages[i].ReceiverUsername == user2) || (messages[i].SenderUsername == user2 && messages[i].ReceiverUsername == user1) {
			messagesBySenderAndReceiver = append(messagesBySenderAndReceiver, messages[i])
		}
	}

	return messagesBySenderAndReceiver, nil
}

func (store *MessageMongoDBStore) GetMessagesByUsername(username string) ([]*domain.Message, error) {
	filter := bson.M{ "$or": bson.A{
						bson.M{"senderUsername": username},
						bson.M{"receiverUsername": username},
					}, }
	return store.filter(filter)
}

func (store *MessageMongoDBStore) Insert(message *domain.Message) (string, error) {
	message.Id = primitive.NewObjectID()
	messageInDatabase, err := store.Get(message.Id)
	if messageInDatabase != nil {
		return "id exists", nil
	}
	result, err := store.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return "error while inserting", err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *MessageMongoDBStore) DeleteAll() {
	store.messages.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *MessageMongoDBStore) filter(filter interface{}) ([]*domain.Message, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessageMongoDBStore) filterOne(filter interface{}) (message *domain.Message, err error) {
	result := store.messages.FindOne(context.TODO(), filter)
	err = result.Decode(&message)
	return
}

func decode(cursor *mongo.Cursor) (messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var message domain.Message
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}
	err = cursor.Err()
	return
}
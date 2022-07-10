package persistence

import (
	"context"
	"time"

	"strings"

	"github.com/XWS-2022-Tim12/Dislinkt/back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notification"
	COLLECTION = "notification"
)

type NotificationMongoDBStore struct {
	notifications *mongo.Collection
}

func NewNotificationMongoDBStore(client *mongo.Client) domain.NotificationStore {
	notifications := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationMongoDBStore{
		notifications: notifications,
	}
}

func (store *NotificationMongoDBStore) Get(id primitive.ObjectID) (*domain.Notification, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *NotificationMongoDBStore) GetAll() ([]*domain.Notification, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *NotificationMongoDBStore) SearchBySender(content string) ([]*domain.Notification, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Notification{}
	allNotifications, _ := store.filter(filter)
	for _, notificationInDatabase := range allNotifications {
		if strings.Contains(notificationInDatabase.Sender, content) {
			contentToSend = append(contentToSend, notificationInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *NotificationMongoDBStore) SearchByReceiver(content string) ([]*domain.Notification, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Notification{}
	allNotifications, _ := store.filter(filter)
	for _, notificationInDatabase := range allNotifications {
		if strings.Contains(notificationInDatabase.Receiver, content) {
			contentToSend = append(contentToSend, notificationInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *NotificationMongoDBStore) SearchByNotificationType(content string) ([]*domain.Notification, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Notification{}
	allNotifications, _ := store.filter(filter)
	for _, notificationInDatabase := range allNotifications {
		if strings.Contains(notificationInDatabase.NotificationType, content) {
			contentToSend = append(contentToSend, notificationInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *NotificationMongoDBStore) Insert(notification *domain.Notification) (string, error) {
	notification.Id = primitive.NewObjectID()
	notification.CreationDate = time.Now()
	notification.IsRead = false
	notificationInDatabase, err := store.Get(notification.Id)
	if notificationInDatabase != nil {
		return "id exists", nil
	}

	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return "error while inserting", err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *NotificationMongoDBStore) DeleteAll() {
	store.notifications.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *NotificationMongoDBStore) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *NotificationMongoDBStore) filterOne(filter interface{}) (notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&notification)
	return
}

func decode(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (store *NotificationMongoDBStore) Edit(notification *domain.Notification) (string, error) {
	notificationFromDatabase, err := store.Get(notification.Id)
	if notificationFromDatabase == nil {
		return "notification doesn't exist", nil
	}
	
	notificationFromDatabase.IsRead = notification.IsRead

	filter := bson.M{"_id": notification.Id}
	update := bson.M{
		"$set": notificationFromDatabase,
	}
	_, err = store.notifications.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil

}

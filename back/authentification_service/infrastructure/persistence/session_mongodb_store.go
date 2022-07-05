package persistence

import (
	"context"
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "authentification"
	COLLECTION = "session"
)

type SessionMongoDBStore struct {
	sessions *mongo.Collection
}

func NewSessionMongoDBStore(client *mongo.Client) domain.SessionStore {
	sessions := client.Database(DATABASE).Collection(COLLECTION)
	return &SessionMongoDBStore{
		sessions: sessions,
	}
}

func (store *SessionMongoDBStore) Get(id primitive.ObjectID) (*domain.Session, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *SessionMongoDBStore) GetAll() ([]*domain.Session, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *SessionMongoDBStore) Insert(session *domain.Session) (string, error) {
	sessionsForDelete, err := store.GetAll()
	for _, sessionForDelete := range sessionsForDelete {
		if sessionForDelete.Date.Add(time.Minute * 30).Before(time.Now()) {
			store.Delete(sessionForDelete.Id)
		}
	}

	sessionInDatabase, err := store.Get(session.Id)
	if sessionInDatabase != nil {
		filterId := bson.M{"_id": sessionInDatabase.Id}
		store.sessions.DeleteOne(context.TODO(), filterId)
		sessionInDatabase.Id = primitive.NewObjectID()
		sessionInDatabase.Date = time.Now()
		sessionInDatabase.UserId = session.UserId
		_, err = store.sessions.InsertOne(context.TODO(), sessionInDatabase)
		if err != nil {
			return "error while inserting", err
		}
		return sessionInDatabase.Id.Hex(), nil
	}

	filter := bson.M{"userId": session.UserId}
	sessionInDatabase, err = store.filterOne(filter)

	if sessionInDatabase != nil {
		store.sessions.DeleteOne(context.TODO(), filter)
		sessionInDatabase.Id = primitive.NewObjectID()
		sessionInDatabase.Date = time.Now()
		_, err = store.sessions.InsertOne(context.TODO(), sessionInDatabase)
		if err != nil {
			return "error while inserting", err
		}
		return sessionInDatabase.Id.Hex(), nil
	}

	session.Id = primitive.NewObjectID()
	session.Date = time.Now()
	result, err := store.sessions.InsertOne(context.TODO(), session)
	if err != nil {
		return "error while inserting", err
	}
	session.Id = result.InsertedID.(primitive.ObjectID)
	return session.Id.Hex(), nil
}

func (store *SessionMongoDBStore) DeleteAll() {
	store.sessions.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *SessionMongoDBStore) Delete(id primitive.ObjectID) {
	filter := bson.M{"_id": id}
	store.sessions.DeleteOne(context.TODO(), filter)
}

func (store *SessionMongoDBStore) DeleteByUserId(userId string) error {
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_userId": id}
	_, err := store.sessions.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (store *SessionMongoDBStore) filter(filter interface{}) ([]*domain.Session, error) {
	cursor, err := store.sessions.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *SessionMongoDBStore) filterOne(filter interface{}) (session *domain.Session, err error) {
	result := store.sessions.FindOne(context.TODO(), filter)
	err = result.Decode(&session)
	return
}

func decode(cursor *mongo.Cursor) (sessions []*domain.Session, err error) {
	for cursor.Next(context.TODO()) {
		var session domain.Session
		err = cursor.Decode(&session)
		if err != nil {
			return
		}
		sessions = append(sessions, &session)
	}
	err = cursor.Err()
	return
}

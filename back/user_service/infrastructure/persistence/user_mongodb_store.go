package persistence

import (
	"context"

	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) UpdateBasicInfo(user *domain.User) (string, error) {
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "user doesn't exist", nil
	}
	if userInDatabase.Password != user.Password {
		return "wrong password", nil
	}
	checkUsername, err := store.GetByUsername(user.Username)
	if checkUsername != nil {
		if checkUsername.Id != userInDatabase.Id {
			return "username exists", nil
		}
	}
	userInDatabase.Firstname = user.Firstname
	userInDatabase.Email = user.Email
	userInDatabase.MobileNumber = user.MobileNumber
	userInDatabase.Gender = user.Gender
	userInDatabase.BirthDay = user.BirthDay
	userInDatabase.Username = user.Username
	userInDatabase.Biography = user.Biography
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil

}
func (store *UserMongoDBStore) UpdateAdvancedInfo(user *domain.User) (string, error) {
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "user doesn't exist", nil
	}
	if userInDatabase.Password != user.Password {
		return "wrong password", nil
	}
	userInDatabase.Experience = user.Experience
	userInDatabase.Education = user.Education
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil

}
func (store *UserMongoDBStore) UpdatePersonalInfo(user *domain.User) (string, error) {

	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "user doesn't exist", nil
	}
	if userInDatabase.Password != user.Password {
		return "wrong password", nil
	}
	userInDatabase.Skills = user.Skills
	userInDatabase.Interests = user.Interests
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil
}
func (store *UserMongoDBStore) UpdateAllInfo(user *domain.User) (string, error) {
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "user doesn't exist", nil
	}
	if userInDatabase.Password != user.Password {
		return "wrong password", nil
	}
	checkUsername, err := store.GetByUsername(user.Username)
	if checkUsername != nil {
		if checkUsername.Id != userInDatabase.Id {
			return "username exists", nil
		}
	}
	userInDatabase.Firstname = user.Firstname
	userInDatabase.Email = user.Email
	userInDatabase.MobileNumber = user.MobileNumber
	userInDatabase.Gender = user.Gender
	userInDatabase.BirthDay = user.BirthDay
	userInDatabase.Username = user.Username
	userInDatabase.Biography = user.Biography
	userInDatabase.Experience = user.Experience
	userInDatabase.Education = user.Education
	userInDatabase.Skills = user.Skills
	userInDatabase.Interests = user.Interests
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil

}

func (store *UserMongoDBStore) Insert(user *domain.User) (string, error) {
	user.Id = primitive.NewObjectID()
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase != nil {
		return "id exists", nil
	}
	userInDatabase, err = store.GetByEmail(user.Email)
	if userInDatabase != nil {
		return "email exists", nil
	}
	userInDatabase, err = store.GetByUsername(user.Username)
	if userInDatabase != nil {
		return "username exists", nil
	}
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return "error while inserting", err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}

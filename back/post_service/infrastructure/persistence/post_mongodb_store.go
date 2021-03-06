package persistence

import (
	"context"
	"sort"

	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "post"
)

type PostMongoDBStore struct {
	posts *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{
		posts: posts,
	}
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetUserPosts(username string) ([]*domain.Post, error) {
	filter := bson.M{"username": username}
	userPosts, err := store.filter(filter)
	if err != nil {
		return nil, err
	}

	sort.Slice(userPosts, func(i, j int) bool { return userPosts[i].Date.After(userPosts[j].Date) })

	return userPosts, nil
}

func (store *PostMongoDBStore) Insert(post *domain.Post) (string, error) {
	post.Id = primitive.NewObjectID()
	postInDatabase, err := store.Get(post.Id)
	if postInDatabase != nil {
		return "id exists", nil
	}
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return "error while inserting", err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *PostMongoDBStore) DeleteAll() {
	store.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}

func (store *PostMongoDBStore) UpdatePost(post *domain.Post) (string, error) {
	postFromDatabase, err := store.Get(post.Id)
	if postFromDatabase == nil {
		return "post doesn't exist", nil
	}
	postFromDatabase.Text = post.Text
	postFromDatabase.Likes = post.Likes
	postFromDatabase.Dislikes = post.Dislikes
	postFromDatabase.Comments = post.Comments
	postFromDatabase.Username = post.Username

	filter := bson.M{"_id": post.Id}
	update := bson.M{
		"$set": postFromDatabase,
	}
	_, err = store.posts.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "error while updating", err
	}

	return "success", nil

}

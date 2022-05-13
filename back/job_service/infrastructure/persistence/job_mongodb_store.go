package persistence

import (
	"context"
	"time"

	"strings"

	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "job"
	COLLECTION = "job"
)

type JobMongoDBStore struct {
	jobs *mongo.Collection
}

func NewJobMongoDBStore(client *mongo.Client) domain.JobStore {
	jobs := client.Database(DATABASE).Collection(COLLECTION)
	return &JobMongoDBStore{
		jobs: jobs,
	}
}

func (store *JobMongoDBStore) Get(id primitive.ObjectID) (*domain.Job, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *JobMongoDBStore) SearchByUser(id primitive.ObjectID) ([]*domain.Job, error) {
	filter := bson.M{"userId": id}
	return store.filter(filter)
}

func (store *JobMongoDBStore) GetAll() ([]*domain.Job, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *JobMongoDBStore) SearchByDescription(content string) ([]*domain.Job, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Job{}
	allJobs, _ := store.filter(filter)
	for _, jobInDatabase := range allJobs {
		if strings.Contains(jobInDatabase.Description, content) {
			contentToSend = append(contentToSend, jobInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *JobMongoDBStore) SearchByPosition(content string) ([]*domain.Job, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Job{}
	allJobs, _ := store.filter(filter)
	for _, jobInDatabase := range allJobs {
		if strings.Contains(jobInDatabase.Position, content) {
			contentToSend = append(contentToSend, jobInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *JobMongoDBStore) SearchByRequirements(content string) ([]*domain.Job, error) {
	filter := bson.D{{}}
	contentToSend := []*domain.Job{}
	allJobs, _ := store.filter(filter)
	for _, jobInDatabase := range allJobs {
		if strings.Contains(jobInDatabase.Requirements, content) {
			contentToSend = append(contentToSend, jobInDatabase)
		}
	}
	return contentToSend, nil
}

func (store *JobMongoDBStore) Insert(job *domain.Job) (string, error) {
	job.Id = primitive.NewObjectID()
	job.CreationDay = time.Now()
	jobInDatabase, err := store.Get(job.Id)
	if jobInDatabase != nil {
		return "id exists", nil
	}
	result, err := store.jobs.InsertOne(context.TODO(), job)
	if err != nil {
		return "error while inserting", err
	}
	job.Id = result.InsertedID.(primitive.ObjectID)
	return "success", nil
}

func (store *JobMongoDBStore) DeleteAll() {
	store.jobs.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *JobMongoDBStore) filter(filter interface{}) ([]*domain.Job, error) {
	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *JobMongoDBStore) filterOne(filter interface{}) (job *domain.Job, err error) {
	result := store.jobs.FindOne(context.TODO(), filter)
	err = result.Decode(&job)
	return
}

func decode(cursor *mongo.Cursor) (jobs []*domain.Job, err error) {
	for cursor.Next(context.TODO()) {
		var job domain.Job
		err = cursor.Decode(&job)
		if err != nil {
			return
		}
		jobs = append(jobs, &job)
	}
	err = cursor.Err()
	return
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

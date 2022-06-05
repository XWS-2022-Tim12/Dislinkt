package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobStore interface {
	Get(id primitive.ObjectID) (*Job, error)
	SearchByUser(id primitive.ObjectID) ([]*Job, error)
	GetAll() ([]*Job, error)
	SearchByDescription(content string) ([]*Job, error)
	SearchByPosition(content string) ([]*Job, error)
	SearchByRequirements(content string) ([]*Job, error)
	Insert(job *Job) (string, error)
	Edit(job *Job) (string, error)
	DeleteAll()
}

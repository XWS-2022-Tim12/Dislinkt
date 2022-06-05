package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	Id                 primitive.ObjectID `bson:"_id"`
	UserId             primitive.ObjectID `bson:"userId"`
	CreationDay        time.Time          `bson:"creationDay"`
	Position           string             `bson:"position"`
	Description        string             `bson:"description"`
	Requirements       string             `bson:"requirements"`
	Comments           []string           `bson:"comments"`
	JuniorSalary       []int32            `bson:"juniorSalary"`
	MediorSalary       []int32            `bson:"mediorSalary"`
	HrInterviews       []string           `bson:"hrInterviews"`
	TehnicalInterviews []string           `bson:"tehnicalInterviews"`
}

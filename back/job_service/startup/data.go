package startup

import (
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobs = []*domain.Job{
	{
		Id:                 getObjectId("62fghcc3a34d25d8567f9f22"),
		UserId:             getObjectId("629d15e1cdfebd707ce8faec"),
		Position:           "Director",
		Description:        "Description of job",
		Requirements:       "Requirements of job",
		CreationDay:        time.Now(),
		Comments:           []string{},
		JuniorSalary:       []int32{},
		MediorSalary:       []int32{},
		HrInterviews:       []string{},
		TehnicalInterviews: []string{},
	},
	{
		Id:                 getObjectId("62fghcc3a34d25d8567f9f23"),
		UserId:             getObjectId("62fsfag3a34d25d8567f9f83"),
		Position:           "Worker",
		Description:        "Description of job",
		Requirements:       "Requirements of job",
		CreationDay:        time.Now(),
		Comments:           []string{},
		JuniorSalary:       []int32{},
		MediorSalary:       []int32{},
		HrInterviews:       []string{},
		TehnicalInterviews: []string{},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

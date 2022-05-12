package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapJob(job *domain.Job) *pb.Job {
	jobPb := &pb.Job{
		Id:           job.Id.Hex(),
		UserId:       job.UserId.Hex(),
		Requirements: job.Requirements,
		Description:  job.Description,
		Position:     job.Position,
		CreationDay:  timestamppb.New(job.CreationDay),
	}
	return jobPb
}

func mapNewJob(jobPb *pb.Job) *domain.Job {
	id, _ := primitive.ObjectIDFromHex(jobPb.UserId)
	job := &domain.Job{
		Id:           primitive.NewObjectID(),
		UserId:       id,
		Requirements: jobPb.Requirements,
		Description:  jobPb.Description,
		Position:     jobPb.Position,
		CreationDay:  time.Now(),
	}
	return job
}

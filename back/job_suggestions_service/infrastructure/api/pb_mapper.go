package api

import (
	"strconv"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
)

func mapJob(jobPb *domain.Job) *pb.Job {
	x := int(jobPb.Id)
	job := &pb.Job{
		Id:           strconv.Itoa(x),
		Position:     jobPb.Position,
		Description:  jobPb.Description,
		Requirements: jobPb.Requirements,
	}
	return job
}

func mapNewJob(jobPb *pb.Job) *domain.Job {
	job := &domain.Job{
		Position:     jobPb.Position,
		Description:  jobPb.Description,
		Requirements: jobPb.Requirements,
	}
	return job
}

package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/application"
)

type JobSuggestionsHandler struct {
	pb.UnimplementedJobSuggestionsServiceServer
	service *application.JobSuggestionsService
}

func NewJobSuggestionsHandler(service *application.JobSuggestionsService) *JobSuggestionsHandler {
	return &JobSuggestionsHandler{
		service: service,
	}
}

func (handler *JobSuggestionsHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	job := mapNewJob(request.Job)
	_, err := handler.service.Register(job)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (handler *JobSuggestionsHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	Suggestions, _ := handler.service.GetAll()

	response := &pb.GetAllResponse{}

	for _, suggestion := range Suggestions {
		current := mapJob(suggestion)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

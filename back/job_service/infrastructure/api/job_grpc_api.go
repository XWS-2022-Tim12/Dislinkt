package api

import (
	"context"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobHandler struct {
	pb.UnimplementedJobServiceServer
	service *application.JobService
}

func NewJobHandler(service *application.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}

func (handler *JobHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	job, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	jobPb := mapJob(job)
	response := &pb.GetResponse{
		Job: jobPb,
	}
	return response, nil
}

func (handler *JobHandler) SearchByUser(ctx context.Context, request *pb.SearchByUserRequest) (*pb.SearchByUserResponse, error) {
	id := request.UserId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	jobs, err := handler.service.SearchByUser(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByUserResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	jobs, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) SearchByDescription(ctx context.Context, request *pb.SearchByDescriptionRequest) (*pb.SearchByDescriptionResponse, error) {
	content := request.Description
	jobs, err := handler.service.SearchByDescription(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByDescriptionResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) SearchByPosition(ctx context.Context, request *pb.SearchByPositionRequest) (*pb.SearchByPositionResponse, error) {
	content := request.Position
	jobs, err := handler.service.SearchByPosition(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByPositionResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) SearchByRequirements(ctx context.Context, request *pb.SearchByRequirementsRequest) (*pb.SearchByRequirementsResponse, error) {
	content := request.Requirements
	jobs, err := handler.service.SearchByRequirements(content)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchByRequirementsResponse{
		Jobs: []*pb.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	job := mapNewJob(request.Job)
	successs, err := handler.service.Add(job)
	response := &pb.AddResponse{
		Success: successs,
	}
	return response, err
}

package api

import (
	"context"
	"log"
	"os"

	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/tracer"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
	otgo "github.com/opentracing/opentracing-go"
)

var (
    InfoLogger  *log.Logger
	ErrorLogger *log.Logger
    trace       otgo.Tracer
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

func init() {
    trace, _ = tracer.Init("job-service")
    otgo.SetGlobalTracer(trace)
    infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

    errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
    if err1 != nil {
        log.Fatal(err1)
    }
    ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *JobHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	job, err := handler.service.Get(objectId)
	if err != nil {
		ErrorLogger.Println("Action: 7, Message: Can not retrieve job!")
		return nil, err
	}
	InfoLogger.Println("Action: 8, Message: Job retrieved successfully!")
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
	if err != nil {
		ErrorLogger.Println("Action: 3, Message: Can not add job!")
        return nil, err
	}
	InfoLogger.Println("Action: 4, Message: Job added successfully!")
	response := &pb.AddResponse{
		Success: successs,
	}
	return response, err
}

func (handler *JobHandler) Edit(ctx context.Context, request *pb.EditRequest) (*pb.EditResponse, error) {
	job := mapChangesOfJob(request.Job)
	successs, err := handler.service.Edit(job)
	if err != nil {
		ErrorLogger.Println("Action: 5, Message: Can not edit job!")
        return nil, err
	}
	InfoLogger.Println("Action: 6, Message: Job edited successfully!")
	response := &pb.EditResponse{
		Success: successs,
	}
	return response, err
}

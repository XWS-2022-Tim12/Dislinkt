package api

import (
	"context"
	"log"
	"os"

	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/tracer"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/application"
	otgo "github.com/opentracing/opentracing-go"
)

var (
    InfoLogger  *log.Logger
	ErrorLogger *log.Logger
    trace       otgo.Tracer
)

type JobSuggestionsHandler struct {
	pb.UnimplementedJobSuggestionsServiceServer
	service *application.JobSuggestionsService
}

func init() {
    trace, _ = tracer.Init("job-suggestions-service")
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

func NewJobSuggestionsHandler(service *application.JobSuggestionsService) *JobSuggestionsHandler {
	return &JobSuggestionsHandler{
		service: service,
	}
}

func (handler *JobSuggestionsHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	job := mapNewJob(request.Job)
	_, err := handler.service.Register(job)
	if err != nil {
		ErrorLogger.Println("Action: 9, Message: Can not register job suggestion!")
		return nil, err
	}
	InfoLogger.Println("Action: 10, Message: Job suggestion registered successfully!")
	return &pb.RegisterResponse{}, nil
}

func (handler *JobSuggestionsHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	Suggestions, err := handler.service.GetAll()
	if err != nil {
		ErrorLogger.Println("Action: 11, Message: Can not retrieve job suggestions!")
		return nil, err
	}
	InfoLogger.Println("Action: 12, Message: Job suggestions retrieved successfully!")

	response := &pb.GetAllResponse{}

	for _, suggestion := range Suggestions {
		current := mapJob(suggestion)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

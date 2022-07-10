package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
)

type JobSuggestionsService struct {
	graph domain.JobSuggestionsGraph
}

func NewJobSuggestionsService(graph domain.JobSuggestionsGraph) *JobSuggestionsService {
	return &JobSuggestionsService{
		graph: graph,
	}
}

func (service *JobSuggestionsService) GetAll() ([]*domain.Job, error) {
	return service.graph.GetAll()
}

func (service *JobSuggestionsService) Register(job *domain.Job) (int64, error) {
	return service.graph.Register(job)
}

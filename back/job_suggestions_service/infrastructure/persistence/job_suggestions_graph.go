package persistence

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type JobSuggestionsDBGraph struct {
	session *neo4j.Session
}

func NewJobSuggestionsGraph(session *neo4j.Session) domain.JobSuggestionsGraph {
	return &JobSuggestionsDBGraph{
		session: session,
	}
}

func (store *JobSuggestionsDBGraph) Register(job *domain.Job) (int64, error) {
	var session = *store.session
	jobId, err := Register(session, job)

	return jobId, err
}

func (store *JobSuggestionsDBGraph) GetAll() ([]*domain.Job, error) {
	var session = *store.session
	jobs, err := GetAll(session)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (store *JobSuggestionsDBGraph) DeleteAll() {
	var session = *store.session
	DeleteAll(session)
	return
}

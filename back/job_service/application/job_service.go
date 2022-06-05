package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	store domain.JobStore
}

func NewJobService(store domain.JobStore) *JobService {
	return &JobService{
		store: store,
	}
}

func (service *JobService) Get(id primitive.ObjectID) (*domain.Job, error) {
	return service.store.Get(id)
}

func (service *JobService) SearchByUser(id primitive.ObjectID) ([]*domain.Job, error) {
	return service.store.SearchByUser(id)
}

func (service *JobService) GetAll() ([]*domain.Job, error) {
	return service.store.GetAll()
}

func (service *JobService) SearchByDescription(content string) ([]*domain.Job, error) {
	return service.store.SearchByDescription(content)
}

func (service *JobService) SearchByPosition(content string) ([]*domain.Job, error) {
	return service.store.SearchByPosition(content)
}

func (service *JobService) SearchByRequirements(content string) ([]*domain.Job, error) {
	return service.store.SearchByRequirements(content)
}

func (service *JobService) Add(job *domain.Job) (string, error) {
	success, err := service.store.Insert(job)
	return success, err
}

func (service *JobService) Edit(job *domain.Job) (string, error) {
	success, err := service.store.Edit(job)
	return success, err
}

package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain"
)

type UserSuggestionsService struct {
	graph domain.UserSuggestionsGraph
}

func NewUserSuggestionsService(graph domain.UserSuggestionsGraph) *UserSuggestionsService {
	return &UserSuggestionsService{
		graph: graph,
	}
}

func (service *UserSuggestionsService) GetAll() ([]*domain.User, error) {
	return service.graph.GetAll()
}

func (service *UserSuggestionsService) Register(user *domain.User) (int64, error) {
	return service.graph.Register(user)
}

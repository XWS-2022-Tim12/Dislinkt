package persistence

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserSuggestionsDBGraph struct {
	session *neo4j.Session
}

func NewUserSuggestionsGraph(session *neo4j.Session) domain.UserSuggestionsGraph {
	return &UserSuggestionsDBGraph{
		session: session,
	}
}

func (store *UserSuggestionsDBGraph) Register(user *domain.User) (int64, error) {
	var session = *store.session

	userId, err := Register(session, user)

	return userId, err
}

func (store *UserSuggestionsDBGraph) GetAll() ([]*domain.User, error) {
	var session = *store.session
	users, err := GetAll(session)
	if err != nil {
		return nil, err
	}
	return users, nil
}

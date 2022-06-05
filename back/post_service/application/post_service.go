package application

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) GetUserPosts(username string) ([]*domain.Post, error) {
	return service.store.GetUserPosts(username)
}

func (service *PostService) AddNewPost(post *domain.Post) (string, error) {
	success, err := service.store.Insert(post)
	return success, err
}

func (service *PostService) LikePost(post *domain.Post) (string, error) {
	success, err := service.store.UpdatePost(post)
	return success, err
}

func (service *PostService) DislikePost(post *domain.Post) (string, error) {
	success, err := service.store.UpdatePost(post)
	return success, err
}

func (service *PostService) CommentPost(post *domain.Post) (string, error) {
	success, err := service.store.UpdatePost(post)
	return success, err
}

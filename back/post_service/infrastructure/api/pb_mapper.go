package api

import (
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Text:     post.Text,
		Image:    post.Image,
		Link:     post.Link,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Comments: post.Comments,
		Username: post.Username,
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {

	post := &domain.Post{
		Id:       primitive.NewObjectID(),
		Text:     postPb.Text,
		Image:    postPb.Image,
		Link:     postPb.Link,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
		Comments: postPb.Comments,
		Username: postPb.Username,
	}
	return post
}

func mapOneMoreLikeToUser(postPb *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)

	post := &domain.Post{
		Id:       id,
		Text:     postPb.Text,
		Image:    postPb.Image,
		Link:     postPb.Link,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
		Comments: postPb.Comments,
		Username: postPb.Username,
	}
	return post
}

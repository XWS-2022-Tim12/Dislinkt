package api

import (
	"time"

	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Text:     post.Text,
		Date:	  timestamppb.New(post.Date),
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Comments: post.Comments,
		Username: post.Username,
		ImageContent: post.ImageContent,
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {

	if postPb.Date != nil {
		post := &domain.Post{
			Id:       primitive.NewObjectID(),
			Text:     postPb.Text,
			Date: 	  postPb.Date.AsTime(),
			Likes:    postPb.Likes,
			Dislikes: postPb.Dislikes,
			Comments: postPb.Comments,
			Username: postPb.Username,
			ImageContent: postPb.ImageContent,
		}
		return post
	} else {
		post := &domain.Post{
			Id:       primitive.NewObjectID(),
			Text:     postPb.Text,
			Date: 	  time.Now(),
			Likes:    postPb.Likes,
			Dislikes: postPb.Dislikes,
			Comments: postPb.Comments,
			Username: postPb.Username,
			ImageContent: postPb.ImageContent,
		}
		return post
	}
}

func mapChangesOfPost(postPb *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)

	if postPb.Date != nil {
		post := &domain.Post{
			Id:       id,
			Text:     postPb.Text,
			Date: 	  postPb.Date.AsTime(),
			Likes:    postPb.Likes,
			Dislikes: postPb.Dislikes,
			Comments: postPb.Comments,
			Username: postPb.Username,
		}
		return post
	} else {
		post := &domain.Post{
			Id:       id,
			Text:     postPb.Text,
			Date: 	  time.Now(),
			Likes:    postPb.Likes,
			Dislikes: postPb.Dislikes,
			Comments: postPb.Comments,
			Username: postPb.Username,
		}
		return post
	}
}

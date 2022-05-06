package api

import (
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:    post.Id.Hex(),
		Text:  post.Text,
		Image: post.Image,
		Link:  post.Link,
	}
	return postPb
}

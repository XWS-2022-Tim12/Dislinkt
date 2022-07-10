package api

import (
	"context"
	"log"
	"os"

	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/tracer"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/post_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
	otgo "github.com/opentracing/opentracing-go"
)

var (
    InfoLogger  *log.Logger
	ErrorLogger *log.Logger
    trace       otgo.Tracer
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func init() {
    trace, _ = tracer.Init("post-service")
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

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		ErrorLogger.Println("Action: 13, Message: Can not retrieve post!")
		return nil, err
	}
	InfoLogger.Println("Action: 14, Message: Post retrieved successfully!")
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetUserPosts(ctx context.Context, request *pb.GetUserPostsRequest) (*pb.GetUserPostsResponse, error) {
	username := request.Username
	
	posts, err := handler.service.GetUserPosts(username)
	if err != nil {
		return nil, err
	}

	response := &pb.GetUserPostsResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) AddNewPost(ctx context.Context, request *pb.AddNewPostRequest) (*pb.AddNewPostResponse, error) {
	post := mapNewPost(request.Post)
	successs, err := handler.service.AddNewPost(post)
	if err != nil {
		ErrorLogger.Println("Action: 15, Message: Can not add post!")
		return nil, err
	}
	InfoLogger.Println("Action: 16, Message: Post added successfully!")

	response := &pb.AddNewPostResponse{
		Success: successs,
	}
	return response, err
}

func (handler *PostHandler) LikePost(ctx context.Context, request *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	post := mapChangesOfPost(request.Post)
	successs, err := handler.service.LikePost(post)
	if err != nil {
		ErrorLogger.Println("Action: 17, Message: Can not like post!")
		return nil, err
	}
	InfoLogger.Println("Action: 18, Message: Post liked successfully!")

	response := &pb.LikePostResponse{
		Success: successs,
	}
	return response, err
}

func (handler *PostHandler) DislikePost(ctx context.Context, request *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	post := mapChangesOfPost(request.Post)
	successs, err := handler.service.DislikePost(post)
	if err != nil {
		ErrorLogger.Println("Action: 19, Message: Can not dislike post!")
		return nil, err
	}
	InfoLogger.Println("Action: 20, Message: Post disliked successfully!")

	response := &pb.DislikePostResponse{
		Success: successs,
	}
	return response, err
}

func (handler *PostHandler) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	post := mapChangesOfPost(request.Post)
	successs, err := handler.service.CommentPost(post)
	if err != nil {
		ErrorLogger.Println("Action: 21, Message: Can not comment post!")
		return nil, err
	}
	InfoLogger.Println("Action: 22, Message: Post commented successfully!")
	response := &pb.CommentPostResponse{
		Success: successs,
	}
	return response, err
}

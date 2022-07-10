package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/infrastructure/services"
	authentification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	job "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_service"
	jobSuggestions "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service"
	post "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/post_service"
	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	userSuggestion "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthentificationHandler struct {
	authentificationClientAddress string
	userClientAddress             string
	postClientAdress              string
	jobClientAdress               string
	jobSuggestionsClientAdress    string
	userSuggestionClientAddress   string
}

func NewAuthentificationHandler(authentificationClientAddress, userClientAddress, postClientAdress, jobClientAdress, userSuggestionClientAddress, jobSuggestionsClientAdress string) Handler {
	return &AuthentificationHandler{
		authentificationClientAddress: authentificationClientAddress,
		userClientAddress:             userClientAddress,
		postClientAdress:              postClientAdress,
		jobClientAdress:               jobClientAdress,
		jobSuggestionsClientAdress:    jobSuggestionsClientAdress,
		userSuggestionClientAddress:   userSuggestionClientAddress,
	}
}

func (handler *AuthentificationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/login", handler.Login)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/user/loginOwner", handler.LoginOwner)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/info/basic", handler.BasicInfo)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/info/advanced", handler.AdvancedInfo)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/info/personal", handler.PersonalInfo)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/info/all", handler.AllInfo)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/user/register", handler.Register)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/follow", handler.FollowPublicProfile)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/acceptFollowingRequest", handler.AcceptFollowingRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/rejectFollowingRequest", handler.RejectFollowingRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/blockUser", handler.BlockUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/user/post/newPost", handler.AddNewPost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/job", handler.AddNewJob)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/user/jobDislinkt", handler.AddNewJobFromDislinkt)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/user/jobDislinktSearch", handler.SearchJobFromDislinkt)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/post/likePost", handler.LikePost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/post/dislikePost", handler.DislikePost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/user/post/commentPost", handler.CommentPost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/user/post/findPosts", handler.FindPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/user/post/findUserPosts/{username}", handler.FindUserPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/user/suggestions/users", handler.GetSuggestionsForLoggedUser)
	if err != nil {
		panic(err)
	}
}

func (handler *AuthentificationHandler) Register(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)
	userToSend := &user.User{
		Id:           usr.Id,
		Firstname:    usr.Firstname,
		Email:        usr.Email,
		MobileNumber: usr.MobileNumber,
		Gender:       mapGender(usr.Gender),
		BirthDay:     timestamppb.New(usr.BirthDay),
		Username:     usr.Username,
		Biography:    usr.Biography,
		Experience:   usr.Experience,
		Education:    mapEducation(usr.Education),
		Skills:       usr.Skills,
		Interests:    usr.Interests,
		Password:     usr.Password,
		Public:       usr.Public,
	}

	idUser, err := userClient.Register(context.TODO(), &user.RegisterRequest{User: userToSend}) //ovdje vratiti userID
	userSuggestionsClient := services.NewUserSuggestionsClient(handler.userSuggestionClientAddress)
	suggestionToSend := &userSuggestion.User{
		FirstName: usr.Firstname,
		Email:     usr.Email,
		Username:  usr.Username,
		Interests: usr.Interests,
	}
	idSuggestion, err := userSuggestionsClient.Register(context.TODO(), &userSuggestion.RegisterRequest{User: suggestionToSend})
	if idSuggestion == nil {
		return
	}

	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	time.Sleep(4 * time.Second)
	success, err := authentificationClient.GetByUserId(context.TODO(), &authentification.GetByUserIdRequest{UserId: idUser.Success})
	cookie := &http.Cookie{Name: "sessionId", Value: success.Session.Id, HttpOnly: false}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	return
}

func (handler *AuthentificationHandler) AllInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:              usr.Id,
		Firstname:       usr.Firstname,
		Email:           usr.Email,
		MobileNumber:    usr.MobileNumber,
		Gender:          mapGender(usr.Gender),
		BirthDay:        timestamppb.New(usr.BirthDay),
		Username:        usr.Username,
		Biography:       usr.Biography,
		Experience:      usr.Experience,
		Education:       mapEducation(usr.Education),
		Skills:          usr.Skills,
		Interests:       usr.Interests,
		FollowedByUsers: usr.FollowedByUsers,
		Password:        usr.Password,
		Public:          usr.Public,
		BlockedUsers:    usr.BlockedUsers,
	}

	userResponse, err := userClient.UpdateAllInfo(context.TODO(), &user.UpdateAllInfoRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) PersonalInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:        usr.Id,
		Skills:    usr.Skills,
		Interests: usr.Interests,
		Password:  usr.Password,
	}

	userResponse, err := userClient.UpdatePersonalInfo(context.TODO(), &user.UpdatePersonalInfoRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) AdvancedInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:         usr.Id,
		Experience: usr.Experience,
		Education:  mapEducation(usr.Education),
		Password:   usr.Password,
	}

	userResponse, err := userClient.UpdateAdvancedInfo(context.TODO(), &user.UpdateAdvancedInfoRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) BasicInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:           usr.Id,
		Firstname:    usr.Firstname,
		Email:        usr.Email,
		MobileNumber: usr.MobileNumber,
		Gender:       mapGender(usr.Gender),
		BirthDay:     timestamppb.New(usr.BirthDay),
		Username:     usr.Username,
		Biography:    usr.Biography,
		Password:     usr.Password,
	}

	userResponse, err := userClient.UpdateBasicInfo(context.TODO(), &user.UpdateBasicInfoRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) IsUserLoggedIn(id string) (string, error) {
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: id})
	if err != nil {
		return "", err
	}
	if success.Session.Role == "user" {
		return "user", nil
	}
	if success.Session.Role == "agent_owner" {
		return "agent_owner", nil
	}

	return "admin", nil
}

func (handler *AuthentificationHandler) findUsersUsername(id string) (string, error) {
	userClient := services.NewUserClient(handler.userClientAddress)
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: id})

	userResponse, err := userClient.GetAll(context.TODO(), &user.GetAllRequest{})
	if err != nil {
		return "", err
	}

	for _, userInDatabase := range userResponse.Users {
		if success.Session.UserId == userInDatabase.Id {
			return userInDatabase.Username, nil
		}
	}

	return "", nil
}

func (handler *AuthentificationHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := handler.FindUser(usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	success, err := handler.AddSession(id, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Name: "sessionId", Value: success, HttpOnly: false}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	return
}

func (handler *AuthentificationHandler) LoginOwner(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success, err := handler.AddSessionOwner(usr.Id, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(success))
	return
}

func (handler *AuthentificationHandler) FindUser(usr *domain.User) (string, error) {
	userClient := services.NewUserClient(handler.userClientAddress)

	userResponse, err := userClient.GetAll(context.TODO(), &user.GetAllRequest{})
	if err != nil {
		return "", err
	}

	for _, userInDatabase := range userResponse.Users {
		if usr.Username == userInDatabase.Username && usr.Password == userInDatabase.Password {
			return userInDatabase.Id, nil
		}
	}

	return "", nil
}

func (handler *AuthentificationHandler) AddSession(id string, r *http.Request) (string, error) {
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	tokenCookie, err := r.Cookie("sessionId")
	session := &pb.Session{
		Id:     "623b0cc3a34d25d8567f9f89",
		UserId: id,
		Role:   "user",
	}
	if err == nil {
		session.Id = tokenCookie.Value
	}

	success, err := authentificationClient.Add(context.TODO(), &authentification.AddRequest{Session: session})
	if err != nil {
		return "", err
	}

	return success.Success, nil
}

func (handler *AuthentificationHandler) AddSessionOwner(id string, r *http.Request) (string, error) {
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	tokenCookie, err := r.Cookie("sessionId")
	session := &pb.Session{
		Id:     "623b0cc3a34d25d8567f9f89",
		UserId: id,
		Role:   "agent_owner",
	}
	if err == nil {
		session.Id = tokenCookie.Value
	}

	success, err := authentificationClient.Add(context.TODO(), &authentification.AddRequest{Session: session})
	if err != nil {
		return "", err
	}

	return success.Success, nil
}

func mapEducation(status string) user.User_EducationEnum {
	switch status {
	case "Primary education":
		return user.User_PrimaryEducation
	case "Lower secondary education":
		return user.User_LowerSecondaryEducation
	case "Upper secondary education":
		return user.User_UpperSecondaryEducation
	case "Bachelor":
		return user.User_Bachelor
	case "Master":
		return user.User_Master
	}
	return user.User_Doctorate

}

func mapEducationInverse(status user.User_EducationEnum) string {
	switch status {
	case user.User_PrimaryEducation:
		return "Primary education"
	case user.User_LowerSecondaryEducation:
		return "Lower secondary education"
	case user.User_UpperSecondaryEducation:
		return "Upper secondary education"
	case user.User_Bachelor:
		return "Bachelor"
	case user.User_Master:
		return "Master"
	}
	return "Doctorate"

}

func mapGender(status string) user.User_GenderEnum {
	switch status {
	case "Male":
		return user.User_Male
	}
	return user.User_Female

}

func (handler *AuthentificationHandler) FollowPublicProfile(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:       usr.Id,
		Username: usr.Username,
	}

	userResponse, err := userClient.FollowPublicProfile(context.TODO(), &user.FollowPublicProfileRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) AcceptFollowingRequest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:       usr.Id,
		Username: usr.Username,
	}

	userResponse, err := userClient.AcceptFollowingRequest(context.TODO(), &user.AcceptFollowingRequestRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) RejectFollowingRequest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:       usr.Id,
		Username: usr.Username,
	}

	userResponse, err := userClient.RejectFollowingRequest(context.TODO(), &user.RejectFollowingRequestRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) BlockUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	usr := &domain.User{}
	errr := json.NewDecoder(r.Body).Decode(&usr)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userClient := services.NewUserClient(handler.userClientAddress)

	userToSend := &user.User{
		Id:             usr.Id,
		Username:       usr.Username,
		FollowingUsers: usr.FollowingUsers,
		BlockedUsers:   usr.BlockedUsers,
	}

	userResponse, err := userClient.BlockUser(context.TODO(), &user.BlockUserRequest{User: userToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userResponse.Success))
	return
}

func (handler *AuthentificationHandler) AddNewPost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqPost := &domain.Post{}
	errr := json.NewDecoder(r.Body).Decode(&reqPost)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	postClient := services.NewPostClient(handler.postClientAdress)
	username, err := handler.findUsersUsername(tokenCookie.Value)
	if username == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	postToSend := &post.Post{
		Id:           reqPost.Id,
		Text:         reqPost.Text,
		Date:         timestamppb.New(time.Now()),
		Likes:        0,
		Dislikes:     0,
		Comments:     []string{},
		Username:     username,
		ImageContent: reqPost.ImageContent,
	}
	postResponse, err := postClient.AddNewPost(context.TODO(), &post.AddNewPostRequest{Post: postToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postResponse.Success))
	return
}

func (handler *AuthentificationHandler) AddNewJobFromDislinkt(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqJob := &domain.Job{}
	errr := json.NewDecoder(r.Body).Decode(&reqJob)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err = handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jobClient := services.NewJobClient(handler.jobClientAdress)
	jobSuggestionsClient := services.NewJobSuggestionsClient(handler.jobSuggestionsClientAdress)

	jobToSend := &job.Job{
		Id:                 reqJob.Id,
		UserId:             reqJob.UserId,
		Position:           reqJob.Position,
		Description:        reqJob.Description,
		Requirements:       reqJob.Requirements,
		Comments:           reqJob.Comments,
		JuniorSalary:       reqJob.JuniorSalary,
		MediorSalary:       reqJob.MediorSalary,
		HrInterviews:       reqJob.HrInterviews,
		TehnicalInterviews: reqJob.TehnicalInterviews,
	}
	jobToSendSugg := &jobSuggestions.Job{
		Id:           reqJob.Id,
		UserId:       reqJob.UserId,
		Position:     reqJob.Position,
		Description:  reqJob.Description,
		Requirements: reqJob.Requirements,
	}
	jobResponse, err := jobClient.Add(context.TODO(), &job.AddRequest{Job: jobToSend})
	_, err = jobSuggestionsClient.Register(context.TODO(), &jobSuggestions.RegisterRequest{Job: jobToSendSugg})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jobResponse.Success))
	return
}

func (handler *AuthentificationHandler) SearchJobFromDislinkt(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err = handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: tokenCookie.Value})
	userId := success.Session.UserId
	userClient := services.NewUserClient(handler.userClientAddress)
	user, err := userClient.Get(context.TODO(), &user.GetRequest{Id: userId})
	jobSuggestionsClient := services.NewJobSuggestionsClient(handler.jobSuggestionsClientAdress)
	jobs, err := jobSuggestionsClient.GetAll(context.TODO(), &jobSuggestions.GetAllRequest{})
	newJobs := make([]*jobSuggestions.Job, 0)
	for _, jobInDatabase := range jobs.Jobs {
		if jobInDatabase.UserId != userId {
			if strings.Contains(strings.ToLower(mapEducationInverse(user.User.Education)), strings.ToLower(jobInDatabase.Description)) || strings.Contains(strings.ToLower(mapEducationInverse(user.User.Education)), strings.ToLower(jobInDatabase.Requirements)) || strings.Contains(strings.ToLower(mapEducationInverse(user.User.Education)), strings.ToLower(jobInDatabase.Position)) {
				newJobs = append(newJobs, jobInDatabase)
			} else if strings.Contains(strings.ToLower(user.User.Experience), strings.ToLower(jobInDatabase.Description)) || strings.Contains(strings.ToLower(user.User.Experience), strings.ToLower(jobInDatabase.Requirements)) || strings.Contains(strings.ToLower(user.User.Experience), strings.ToLower(jobInDatabase.Position)) {
				newJobs = append(newJobs, jobInDatabase)
			} else if strings.Contains(strings.ToLower(user.User.Skills), strings.ToLower(jobInDatabase.Description)) || strings.Contains(strings.ToLower(user.User.Skills), strings.ToLower(jobInDatabase.Requirements)) || strings.Contains(strings.ToLower(user.User.Skills), strings.ToLower(jobInDatabase.Position)) {
				newJobs = append(newJobs, jobInDatabase)
			} else if strings.Contains(strings.ToLower(user.User.Interests), strings.ToLower(jobInDatabase.Description)) || strings.Contains(strings.ToLower(user.User.Interests), strings.ToLower(jobInDatabase.Requirements)) || strings.Contains(strings.ToLower(user.User.Interests), strings.ToLower(jobInDatabase.Position)) {
				newJobs = append(newJobs, jobInDatabase)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newJobs)
	return
}

func (handler *AuthentificationHandler) AddNewJob(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqJob := &domain.Job{}
	errr := json.NewDecoder(r.Body).Decode(&reqJob)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("ownerSessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id != "agent_owner" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jobClient := services.NewJobClient(handler.jobClientAdress)

	jobToSend := &job.Job{
		Id:                 reqJob.Id,
		UserId:             reqJob.UserId,
		Position:           reqJob.Position,
		Description:        reqJob.Description,
		Requirements:       reqJob.Requirements,
		Comments:           reqJob.Comments,
		JuniorSalary:       reqJob.JuniorSalary,
		MediorSalary:       reqJob.MediorSalary,
		HrInterviews:       reqJob.HrInterviews,
		TehnicalInterviews: reqJob.TehnicalInterviews,
	}
	jobResponse, err := jobClient.Add(context.TODO(), &job.AddRequest{Job: jobToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jobResponse.Success))
	return
}

func (handler *AuthentificationHandler) isUserFollowing(id string, username string) bool {
	userClient := services.NewUserClient(handler.userClientAddress)
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: id})

	userResponse, err := userClient.GetAll(context.TODO(), &user.GetAllRequest{})
	if err != nil {
		return false
	}

	loggedInUserResponse, err := userClient.Get(context.TODO(), &user.GetRequest{Id: success.Session.UserId})
	if err != nil {
		return false
	}

	if loggedInUserResponse.User.Username == username {
		return true
	}

	for _, userInDatabase := range userResponse.Users {
		if userInDatabase.Username == username {
			if userInDatabase.Public == true {
				return true
			}
		}
	}

	for _, userInDatabase := range userResponse.Users {
		if success.Session.UserId == userInDatabase.Id {
			for _, followingUser := range userInDatabase.FollowingUsers {
				if username == followingUser {
					return true
				}
			}
		}
	}

	return false
}

func (handler *AuthentificationHandler) LikePost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqPost := &domain.Post{}
	errr := json.NewDecoder(r.Body).Decode(&reqPost)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	postClient := services.NewPostClient(handler.postClientAdress)
	retVal := handler.isUserFollowing(tokenCookie.Value, reqPost.Username)
	if retVal == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	postToSend := &post.Post{
		Id:       reqPost.Id,
		Text:     reqPost.Text,
		Likes:    reqPost.Likes + 1,
		Dislikes: reqPost.Dislikes,
		Comments: reqPost.Comments,
		Username: reqPost.Username,
	}

	postResponse, err := postClient.LikePost(context.TODO(), &post.LikePostRequest{Post: postToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postResponse.Success))
	return
}

func (handler *AuthentificationHandler) DislikePost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqPost := &domain.Post{}
	errr := json.NewDecoder(r.Body).Decode(&reqPost)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	postClient := services.NewPostClient(handler.postClientAdress)
	retVal := handler.isUserFollowing(tokenCookie.Value, reqPost.Username)
	if retVal == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	postToSend := &post.Post{
		Id:       reqPost.Id,
		Text:     reqPost.Text,
		Likes:    reqPost.Likes,
		Dislikes: reqPost.Dislikes + 1,
		Comments: reqPost.Comments,
		Username: reqPost.Username,
	}

	postResponse, err := postClient.DislikePost(context.TODO(), &post.DislikePostRequest{Post: postToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postResponse.Success))
	return
}

func (handler *AuthentificationHandler) CommentPost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqPost := &domain.Post{}
	errr := json.NewDecoder(r.Body).Decode(&reqPost)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	postClient := services.NewPostClient(handler.postClientAdress)
	retVal := handler.isUserFollowing(tokenCookie.Value, reqPost.Username)
	if retVal == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	postToSend := &post.Post{
		Id:       reqPost.Id,
		Text:     reqPost.Text,
		Likes:    reqPost.Likes,
		Dislikes: reqPost.Dislikes,
		Comments: reqPost.Comments,
		Username: reqPost.Username,
	}

	postResponse, err := postClient.CommentPost(context.TODO(), &post.CommentPostRequest{Post: postToSend})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postResponse.Success))
	return
}

func (handler *AuthentificationHandler) GetSuggestionsForLoggedUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	tokenCookie, err := r.Cookie("sessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	id, err := handler.IsUserLoggedIn(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: tokenCookie.Value})
	userId := success.Session.UserId
	userClient := services.NewUserClient(handler.userClientAddress)
	user, err := userClient.Get(context.TODO(), &user.GetRequest{Id: userId})

	userSuggestionsClient := services.NewUserSuggestionsClient(handler.userSuggestionClientAddress)
	suggestedUsers := make([]*userSuggestion.User, 0)

	suggestionsResponse, err := userSuggestionsClient.GetAll(context.TODO(), &userSuggestion.GetAllRequest{})
	for _, suggest := range suggestionsResponse.Users {
		if strings.Contains(strings.ToLower(suggest.Interests), strings.ToLower(user.User.Interests)) {
			suggestedUsers = append(suggestedUsers, suggest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suggestedUsers)
	return
}

func (handler *AuthentificationHandler) FindPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userClient := services.NewUserClient(handler.userClientAddress)
	postClient := services.NewPostClient(handler.postClientAdress)

	postsResponse, err := postClient.GetAll(context.TODO(), &post.GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResponse, err := userClient.GetAll(context.TODO(), &user.GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postsByPublicUser := make([]*post.Post, 0)
	for _, postInDatabase := range postsResponse.Posts {
		for _, userInDatabase := range userResponse.Users {
			if postInDatabase.Username == userInDatabase.Username {
				if userInDatabase.Public == true {
					postsByPublicUser = append(postsByPublicUser, postInDatabase)
				} else {
					tokenCookie, err := r.Cookie("sessionId")
					if err != nil {
						break
					}
					id, err := handler.IsUserLoggedIn(tokenCookie.Value)
					if err != nil {
						break
					}
					if id == "" {
						break
					}

					username, err := handler.findUsersUsername(tokenCookie.Value)
					for _, u := range userInDatabase.FollowedByUsers {
						if u == username {
							postsByPublicUser = append(postsByPublicUser, postInDatabase)
							break
						}
					}
				}

				break
			}
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postsByPublicUser)
	return
}

func (handler *AuthentificationHandler) FindUserPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	username := pathParams["username"]
	userClient := services.NewUserClient(handler.userClientAddress)

	userResponse, err := userClient.GetByUsername(context.TODO(), &user.GetByUsernameRequest{Username: username})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userFromResponse := userResponse.User

	if !userFromResponse.Public {
		tokenCookie, err := r.Cookie("sessionId")
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		id, err := handler.IsUserLoggedIn(tokenCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
		success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: tokenCookie.Value})
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		loggedUserResponse, err := userClient.Get(context.TODO(), &user.GetRequest{Id: success.Session.UserId})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		loggedUser := loggedUserResponse.User

		isFollowingUser := false
		for _, followingUsername := range loggedUser.FollowingUsers {
			if followingUsername == username {
				isFollowingUser = true
			}
		}

		if success.Session.UserId == userFromResponse.Id || isFollowingUser {
			postClient := services.NewPostClient(handler.postClientAdress)

			postsResponse, err := postClient.GetUserPosts(context.TODO(), &post.GetUserPostsRequest{Username: username})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var postsToSend []*domain.Post
			for _, post := range postsResponse.Posts {
				current := mapPost(post)
				postsToSend = append(postsToSend, current)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(postsToSend)
			return
		} else {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	} else {
		postClient := services.NewPostClient(handler.postClientAdress)

		postsResponse, err := postClient.GetUserPosts(context.TODO(), &post.GetUserPostsRequest{Username: username})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var postsToSend []*domain.Post
		for _, post := range postsResponse.Posts {
			current := mapPost(post)
			postsToSend = append(postsToSend, current)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postsToSend)
		return
	}
}

func mapPost(postPb *post.Post) *domain.Post {
	if postPb.Date != nil {
		post := &domain.Post{
			Id:           postPb.Id,
			Text:         postPb.Text,
			Date:         postPb.Date.AsTime(),
			Likes:        postPb.Likes,
			Dislikes:     postPb.Dislikes,
			Comments:     postPb.Comments,
			Username:     postPb.Username,
			ImageContent: postPb.ImageContent,
		}
		return post
	} else {
		post := &domain.Post{
			Id:           postPb.Id,
			Text:         postPb.Text,
			Date:         time.Now(),
			Likes:        postPb.Likes,
			Dislikes:     postPb.Dislikes,
			Comments:     postPb.Comments,
			Username:     postPb.Username,
			ImageContent: postPb.ImageContent,
		}
		return post
	}
}

package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/infrastructure/services"
	authentification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthentificationHandler struct {
	authentificationClientAddress string
	userClientAddress             string
	postClientAdress              string
}

func NewAuthentificationHandler(authentificationClientAddress, userClientAddress, postClientAdress string) Handler {
	return &AuthentificationHandler{
		authentificationClientAddress: authentificationClientAddress,
		userClientAddress:             userClientAddress,
		postClientAdress:              postClientAdress,
	}
}

func (handler *AuthentificationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/login", handler.Login)
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
	err = mux.HandlePath("PUT", "/user/follow", handler.FollowPublicProfile)
	if err != nil {
		panic(err)
	}
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

	return "admin", nil
}

func (handler *AuthentificationHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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

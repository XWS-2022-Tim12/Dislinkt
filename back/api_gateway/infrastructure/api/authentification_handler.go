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
)

type AuthentificationHandler struct {
	authentificationClientAddress string
	userClientAddress             string
}

func NewAuthentificationHandler(authentificationClientAddress, userClientAddress string) Handler {
	return &AuthentificationHandler{
		authentificationClientAddress: authentificationClientAddress,
		userClientAddress:             userClientAddress,
	}
}

func (handler *AuthentificationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/login", handler.Login)
	if err != nil {
		panic(err)
	}
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

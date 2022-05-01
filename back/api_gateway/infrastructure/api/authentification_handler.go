package api

import (
	"context"
	"encoding/json"
	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/domain"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/infrastructure/services"
	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service"
	authentification "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	pb "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	"net/http"
)

type AuthentificationHandler struct {
	authentificationClientAddress  string
	userClientAddress string
}

func NewAuthentificationHandler(authentificationClientAddress, userClientAddress string) Handler {
	return &AuthentificationHandler{
		authentificationClientAddress:  authentificationClientAddress,
		userClientAddress: userClientAddress,
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
	userClient := services.NewUserClient(handler.userClientAddress)
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)

	userResponse, err := userClient.GetAll(context.TODO(), &user.GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, userInDatabase := range userResponse.Users {
		if usr.Username == userInDatabase.Username && usr.Password == userInDatabase.Password {
			tokenCookie, err := r.Cookie("sessionId")
			if err != nil {
				tokenCookie = nil
			}

			session := &pb.Session{
				Id:           tokenCookie.Value,
				UserId:       userInDatabase.Id,
				Role:         "user",
			}

			success, err := authentificationClient.Add(context.TODO(), &authentification.AddRequest{Session: session})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			cookie := &http.Cookie{Name: "sessionId", Value: success.Success, HttpOnly: false}
			http.SetCookie(w, cookie)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}
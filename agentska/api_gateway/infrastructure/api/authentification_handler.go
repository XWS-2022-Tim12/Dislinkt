package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/XWS-2022-Tim12/Dislinkt/agentska/api_gateway/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/agentska/api_gateway/infrastructure/services"
	authentification "github.com/XWS-2022-Tim12/Dislinkt/agentska/common/proto/authentification_service"
	pb "github.com/XWS-2022-Tim12/Dislinkt/agentska/common/proto/authentification_service"
	user "github.com/XWS-2022-Tim12/Dislinkt/agentska/common/proto/user_service"
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
	err = mux.HandlePath("POST", "/job", handler.AddNewJob)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/{id}", handler.FindJobById)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/searchByUser/{userId}", handler.FindJobByUserId)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/jobs", handler.FindJobs)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/searchByDescription/{description}", handler.FindJobByDesctiption)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/searchByPosition/{position}", handler.FindJobByPosition)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/job/searchByRequirements/{requirements}", handler.FindJobByRequirements)
	if err != nil {
		panic(err)
	}
}

func (handler *AuthentificationHandler) IsUserLoggedIn(id string) (string, error) {
	authentificationClient := services.NewAuthentificationClient(handler.authentificationClientAddress)
	success, err := authentificationClient.Get(context.TODO(), &authentification.GetRequest{Id: id})
	if err != nil {
		return "", err
	}
	if success.Session.Role == "agent_user" {
		return "agent_user", nil
	}
	if success.Session.Role == "agent_owner" {
		return "agent_owner", nil
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

	cookie := &http.Cookie{Name: "agentSessionId", Value: success, HttpOnly: false, Path: "/job"}
	http.SetCookie(w, cookie)
	role, err := handler.IsUserLoggedIn(success)
	if role == "agent_owner" {

		url := "http://api_gateway:8000/user/loginOwner"

		var jsonStr = []byte(`{"id":"` + usr.Id + `"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		cookiee := &http.Cookie{Name: "ownerSessionId", Value: string(body), HttpOnly: false, Path: "/job"}
		http.SetCookie(w, cookiee)
	}

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
	tokenCookie, err := r.Cookie("agentSessionId")
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

func (handler *AuthentificationHandler) AddNewJob(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reqJob := &domain.Job{}
	errr := json.NewDecoder(r.Body).Decode(&reqJob)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenCookie, err := r.Cookie("agentSessionId")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	ownerTokenCookie, err := r.Cookie("ownerSessionId")
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

	url := "http://api_gateway:8000/job"

	values := map[string]string{"id": reqJob.Id, "userId": reqJob.UserId, "position": reqJob.Position, "description": reqJob.Description, "requirements": reqJob.Requirements}

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	cookie := &http.Cookie{Name: "ownerSessionId", Value: ownerTokenCookie.Value, HttpOnly: false}
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobById(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	content := pathParams["id"]
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job" + "/" + content

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobByUserId(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	content := pathParams["userId"]
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job/searchByUser" + "/" + content

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobs(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job/jobs"

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobByDesctiption(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	content := pathParams["description"]
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job/searchByDescription" + "/" + content

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobByPosition(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	content := pathParams["position"]
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job/searchByPosition" + "/" + content

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

func (handler *AuthentificationHandler) FindJobByRequirements(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	content := pathParams["requirements"]
	tokenCookie, err := r.Cookie("agentSessionId")
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

	url := "http://api_gateway:8000/job/searchByRequirements" + "/" + content

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(body)))

	return
}

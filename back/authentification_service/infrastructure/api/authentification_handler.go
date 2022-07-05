package api

import (
	"fmt"

	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/application"
	saga "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/messaging"
	events "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/register_user"
)

type CreateAuthentificationCommandHandler struct {
	sessionService    *application.SessionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateAuthentificationCommandHandler(sessionService *application.SessionService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateAuthentificationCommandHandler, error) {
	o := &CreateAuthentificationCommandHandler{
		sessionService:    sessionService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateAuthentificationCommandHandler) handle(command *events.RegisterUserCommand) {
	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.AddUserAuthentification:
		fmt.Println("usao u aut handler, dodaj sesiju")
		session := mapCommandToSession(command)
		_, err := handler.sessionService.Add(session)
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			reply.Type = events.UserAuthentificationNotAdded
			break
		}
		//cookie := &http.Cookie{Name: "sessionId", Value: id, HttpOnly: false}
		//http.SetCookie(w, cookie)
		//w.WriteHeader(http.StatusOK)
		reply.Type = events.UserAuthentificationAdded
	case events.RollbackAddUserAuthentification:
		fmt.Println("usao u aut handler, rollback sesiju")
		err := handler.sessionService.DeleteByUserId(command.User.Id)
		if err != nil {
			return
		}
		reply.Type = events.UserAuthentificationRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
	return
}

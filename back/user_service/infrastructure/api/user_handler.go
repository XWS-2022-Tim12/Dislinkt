package api

import (
	"fmt"

	saga "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/messaging"
	events "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/register_user"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateUserCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateUserCommandHandler, error) {
	o := &CreateUserCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateUserCommandHandler) handle(command *events.RegisterUserCommand) {
	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.RollbackAddUser:
		fmt.Println("usao u user handler, rollback user")
		objectId, _ := primitive.ObjectIDFromHex(command.User.Id)
		handler.userService.Delete(objectId)
		reply.Type = events.UserAddRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

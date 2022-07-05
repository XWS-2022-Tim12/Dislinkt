package application

import (
	"fmt"

	saga "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/messaging"
	events "github.com/XWS-2022-Tim12/Dislinkt/back/common/saga/register_user"
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_service/domain"
)

type AddUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewAddUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*AddUserOrchestrator, error) {
	o := &AddUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *AddUserOrchestrator) Start(user *domain.User) error {
	fmt.Println("usao u orkestrator start")
	event := &events.RegisterUserCommand{
		Type: events.AddUserAuthentification,
		User: events.User{
			Id:           user.Id.Hex(),
			Firstname:    user.Firstname,
			Email:        user.Email,
			MobileNumber: user.MobileNumber,
			Gender:       mapNewGender(user.Gender),
			BirthDay:     user.BirthDay,
			Username:     user.Username,
			Biography:    user.Biography,
			Experience:   user.Experience,
			Education:    mapNewEducation(user.Education),
			Skills:       user.Skills,
			Interests:    user.Interests,
			Password:     user.Password,
			Public:       user.Public,
		},
	}
	return o.commandPublisher.Publish(event)
}

func mapNewGender(status domain.GenderEnum) events.GenderEnum {
	switch status {
	case domain.Male:
		return events.Male
	}
	return events.Female

}

func mapNewEducation(status domain.EducationEnum) events.EducationEnum {
	switch status {
	case domain.PrimaryEducation:
		return events.PrimaryEducation
	case domain.LowerSecondaryEducation:
		return events.LowerSecondaryEducation
	case domain.UpperSecondaryEducation:
		return events.UpperSecondaryEducation
	case domain.Bachelor:
		return events.Bachelor
	case domain.Master:
		return events.Master
	}
	return events.Doctorate

}

func (o *AddUserOrchestrator) handle(reply *events.RegisterUserReply) {
	command := events.RegisterUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *AddUserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch reply {
	case events.UserAuthentificationAdded:
		fmt.Println("autentifikacija dodata")
		return events.UnknownCommand
	case events.UserAuthentificationNotAdded:
		fmt.Println("autentifikacija nije dodata")
		return events.RollbackAddUser
	case events.UserAuthentificationRolledBack:
		fmt.Println("rollbekovana autentifikacija")
		return events.RollbackAddUser
	case events.UserAddRolledBack:
		fmt.Println("user obrisan")
		return events.UnknownCommand
	default:
		return events.UnknownCommand
	}
}

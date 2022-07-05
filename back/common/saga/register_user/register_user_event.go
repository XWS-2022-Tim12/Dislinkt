package register_user

import "time"

type User struct {
	Id                string
	Firstname         string
	Email             string
	MobileNumber      string
	Gender            GenderEnum
	BirthDay          time.Time
	Username          string
	Biography         string
	Experience        string
	Education         EducationEnum
	Skills            string
	Interests         string
	Password          string
	FollowingUsers    []string
	FollowedByUsers   []string
	FollowingRequests []string
	Public            bool
}

type EducationEnum int8

type GenderEnum int8

const (
	PrimaryEducation EducationEnum = iota
	LowerSecondaryEducation
	UpperSecondaryEducation
	Bachelor
	Master
	Doctorate
)

const (
	Male GenderEnum = iota
	Female
)

type RegisterUserCommandType int8

const (
	AddUserAuthentification RegisterUserCommandType = iota
	RollbackAddUserAuthentification
	RollbackAddUser
	UnknownCommand
)

type RegisterUserCommand struct {
	User User
	Type RegisterUserCommandType
}

type RegisterUserReplyType int8

const (
	UserAuthentificationAdded RegisterUserReplyType = iota
	UserAuthentificationNotAdded
	UserAuthentificationRolledBack
	UserAddRolledBack
	UnknownReply
)

type RegisterUserReply struct {
	User User
	Type RegisterUserReplyType
}

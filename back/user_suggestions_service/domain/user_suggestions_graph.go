package domain

type UserSuggestionsGraph interface {
	GetAll() ([]*User, error)
	Register(user *User) (int64, error)
	DeleteAll()
}

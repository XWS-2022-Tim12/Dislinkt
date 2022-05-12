package domain

import "time"

type User struct {
	Id           string
	Firstname    string
	Email        string
	MobileNumber string
	Gender       string
	BirthDay     time.Time
	Username     string
	Password     string
}

type Session struct {
	Id     string
	UserId string
	Date   time.Time
	Role   string
}

type Job struct {
	Id           string
	UserId       string
	CreationDay  time.Time
	Position     string
	Description  string
	Requirements string
}

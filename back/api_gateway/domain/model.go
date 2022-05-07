package domain

import "time"

type User struct {
	Id              string
	Firstname       string
	Email           string
	MobileNumber    string
	Gender          string
	BirthDay        time.Time
	Username        string
	Biography       string
	Experience      string
	Education       string
	Skills          string
	Interests       string
	Password        string
	FollowingUsers  []string
	FollowedByUsers []string
}

type Session struct {
	Id     string
	UserId string
	Date   time.Time
	Role   string
}

type Post struct {
	Id       string
	Text     string
	Image    string
	Link     string
	Likes    int32
	Dislikes int32
	Comments []string
	username string
}

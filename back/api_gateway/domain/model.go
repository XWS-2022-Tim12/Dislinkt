package domain

import "time"

type User struct {
	Id                string
	Firstname         string
	Email             string
	MobileNumber      string
	Gender            string
	BirthDay          time.Time
	Username          string
	Biography         string
	Experience        string
	Education         string
	Skills            string
	Interests         string
	Password          string
	FollowingUsers    []string
	FollowedByUsers   []string
	FollowingRequests []string
	Public            bool
	BlockedUsers      []string
	Notifications	  bool
	NotificationOffUsers    []string
	NotificationOffMessages   []string
	Role				string
}

type Session struct {
	Id     string
	UserId string
	Date   time.Time
	Role   string
}

type Post struct {
	Id           string    `json:"id"`
	Text         string    `json:"text"`
	Date         time.Time `json:"date"`
	Likes        int32     `json:"likes"`
	Dislikes     int32     `json:"dislikes"`
	Comments     []string  `json:"comments"`
	Username     string    `json:"username"`
	ImageContent string    `json:"imageContent"`
}

type Message struct {
	Id           string    `json:"id"`
	Text         string    `json:"text"`
	Date         time.Time `json:"date"`
	SenderUsername     string    `json:"senderUsername"`
	ReceiverUsername string    `json:"receiverUsername"`
}

type Job struct {
	Id                 string
	UserId             string
	CreationDay        time.Time
	Position           string
	Description        string
	Requirements       string
	Comments           []string
	JuniorSalary       []int32
	MediorSalary       []int32
	HrInterviews       []string
	TehnicalInterviews []string
}

type Notification struct {
	Id                 string
	Sender             string
	Receiver           string
	CreationDate       time.Time
	NotificationType   string
	Description        string
	IsRead		       bool
}

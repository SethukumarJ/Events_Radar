package model

import (
	"time"

	"gorm.io/gorm"
)

// user schema for user table to get listed all users
type User struct {
	gorm.Model

	Id              int    `json:"user_id"`
	First_Name      string `json:"first_name"`
	Last_Name       string `json:"last_name"`
	Username        string `json:"username" gorm:"primary key"`
	Email           string `json:"email" gorm:"not null;unique"`
	Phone           int64  `json:"phone_number"`
	Password        string `json:"password"`
	Verified        bool   `json:"verified" gorm:"default:false"`
	Verification    bool   `json:"verification" gorm:"default:false"`
	Profile         string `json:"profile"`
	Events_Id       string `json:"events_id"`       //forieng key reference to events
	Notification_Id string `json:"notification_id"` //forieng key reference to notification
	TimeLine_Id     string `json:"timeline_id"`     //forieng key reference to timeline
}

// to store admin credentials
type Admin struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//to store mail verification details

type Verification struct {
	gorm.Model
	Email string `json:"email"`
	Code  int    `json:"code"`
}

// type Bio struct {
// 	Id              int    `json:"user_id"`
// 	Bio_discription string `json:"Bio"`
// 	Linked_in       string `json:"linked in"`
// 	Instagram       string `json:"instagram"`
// 	Github          string `json:"github"`
// 	Facebook        string `json:"facebook"`
// 	User            string `json:"user"` //forieng key reference to user
// }

type Event struct {
	gorm.Model
	Created_at               time.Time `json:"created_at"`
	Organizer_name           string    `json:"orginizer_name"` //forieng key reference to user
	Title                    string    `json:"title" gorm:"not null;unique"`
	Event_date               string    `json:"event_date"`
	Location                 string    `json:"location"`
	Offline                  bool      `json:"offline"`
	Free                     bool      `json:"free"`
	Short_description        string    `json:"short_description"`
	Long_description         string    `json:"long_description"`
	Application_link         string    `json:"application_link"`
	Website_link             string    `json:"website_link"`
	Max_application          int       `json:"max_applications"`
	Sex                      string    `json:"sex"`
	Cusat_only               bool      `json:"cusat_only" gorm:"default:false"`
	Application_closing_date string    `json:"application_closing_date"`
	Sub_events               string    `json:"sub_events"`
	Application_template     string    `json:"application_template"`
	Archived                 bool      `json:"archived"`
	Event_pic                string    `json:"event_pic"`
	Question_id              int       `json:"Question_id"`
	Approved                 bool      `json:"approved" gorm:"default:false"`
}
type ApplicationForm struct {
	gorm.Model
	Event_name     string `json:"event_name"`
	Applicant_name string `json:"name"`
	Organizer_name string `json:"Organizer_name"`
	Proffession    string `json:"proffession"`
	College        string `json:"college"`
	Company        string `json:"company"`
	About          string `json:"about"`
	Email          string `json:"email"`
	Github         string `json:"github"`
	Linkedin       string `json:"linkedin"`
}

type Bookmarks struct {
	gorm.Model
	User_Id  string `json:"user_id"`  //forieng key reference to user
	Event_Id string `json:"event_id"` //forieng key reference to event
}

type Applied_events struct {
	Applied_at           time.Time `json:"applied_at"`
	User_Id              string    `json:"user_id"`  //forieng key reference to user
	Event_Id             string    `json:"event_id"` //forieng key reference to event
	Participation_status bool      `json:"participation_status"`
	Application_accepted bool      `json:"application_accepted" gorm:"defautl:false"`
}

type FAQA struct {
	gorm.Model
	Event_name string `json:"event_name"`
	Question   string `json:"question"`
	Username   string `json:"username"`
	Answer     string `json:"answer"`
	Public     bool   `json:"public" gorm:"default:false"`
	Date       string `json:"date"`
}

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
	Username        string `json:"username"`
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

// type OrdinaryUser struct {
// 	Id                  int    `json:"user_id"`
// 	Proffession         string `json:"Proffesion"`
// 	Participated_events string `json:"participated_events"`
// 	User                string `json:"user"` //forieng key reference to user
// }

// type Organization struct {
// 	Id               int    `json:"user_id"`
// 	Location         string `json:"location"`
// 	Organized_events string `json:"organized_events"`
// 	User             string `json:"user"` //forieng key reference to user
// }

type Event struct {
	ID        				 int  	`json:"id,omitempty"`
	Created_at               time.Time `json:"created_at"`
	Organizer                string `json:"user"` //forieng key reference to user
	Title                    string `json:"title"`
	Event_date               string `json:"event_date"`
	Location                 string `json:"location"`
	Offline                  bool   `json:"offline"`
	Free                     bool   `json:"Free"`
	Short_description        string `json:"short_description"`
	Long_description         string `json:"long_description"`
	Application_link         string `json:"application_link"`
	Website_link             string `json:"website_link"`
	Limit_applications_Id    string `json:"limit_applications"`
	Application_closing_date string `json:"application_closing_date"`
	Sub_events               string `json:"sub_events"`
	Archived                 bool   `json:"archived"`
	Event_pic                string `json:"event_pic"`
	Question_id				 int    `json:"Question_id"`
	Approved				 bool	`json:"approved" gorm:"default:false"`
}





type LimitAppication struct {
	Id					int    `json:"limitapplication_id"`
	Max_application		int    `json:"max_applications"`
	MaleOrFemale		string `json:"sex"`
	CusatOnly			bool   `json:"cusatOnly" gorm:"default:false"`
}

type Bookmarks struct {
	gorm.Model
	User_Id  	string `json:"user_id"`  //forieng key reference to user
	Event_Id 	string `json:"event_id"` //forieng key reference to event
}

type Applied_events struct {
	Applied_at			time.Time `json:"applied_at"`
	User_Id              string `json:"user_id"`  //forieng key reference to user
	Event_Id             string `json:"event_id"` //forieng key reference to event
	Application_status   string `json:"application_status"`
	Participation_status bool   `json:"participation_status"`
	Application_accepted bool 	`json:"application_accepted" gorm:"defautl:false"`
}


package model

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	gorm.Model

	ID         int    `json:"user_id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      int    `json:"phone_number"`
	Password   string `json:"password,omitempty"`
	Verified   bool   `json:"verified" gorm:"default:false"`
	Profile    string `json:"profile"`
	Token      string `json:"token"`
}

type AdminResponse struct {
	ID       int    `json:"id" `
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type EventResponse struct {
	Created_at               time.Time `json:"created_at"`
	Organizer_name           string    `json:"organizer_name"` //forieng key reference to user
	Title                    string    `json:"title"`
	Event_pic                string    `json:"event_pic"`
	Event_date               string    `json:"event_date"`
	Location                 string    `json:"location"`
	Offline                  bool      `json:"offline"`
	Free                     bool      `json:"Free"`
	Short_description        string    `json:"short_description"`
	Long_description         string    `json:"long_description"`
	Application_link         string    `json:"application_link"`
	Website_link             string    `json:"website_link"`
	Application_closing_date string    `json:"application_closing_date"`
	Sub_events               string    `json:"sub_events"`
	Archived                 bool      `json:"archived"`
}


type FAQAResponse struct {
	gorm.Model
	Event_name string `json:"event_name"`
	Question   string `json:"question"`
	User_name  string `json:"user_name"`
	Answer     string `json:"answer"`
	Public     bool	  `json:"public" gorm:"default:false"`
	Date       string `json:"date"`
}

// type BioResponse struct {
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

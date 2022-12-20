package model

import (
	"gorm.io/gorm"
)

// user schema for user table to get listed all users
type User struct {
	gorm.Model

	Id               	int    `json:"user_id"`
	First_Name       	string `json:"first_name"`
	Last_Name        	string `json:"last_name"`
	Email            	string `json:"email" gorm:"not null;unique"`
	Phone            	int64  `json:"phone_number"`
	Password         	string `json:"password"`
	Verified         	bool   `json:"verified" gorm:"default:false"`
	Verification     	bool   `json:"verification" gorm:"default:false"`
	Profile          	string `json:"profile"`
	Events_Id			string `json:"events_id"` //forieng key reference to events
	Notification_Id		string `json:"notification_id"` //forieng key reference to notification
	TimeLine_Id			string `json:"timeline_id"` //forieng key reference to timeline
	
}

type Bio struct {
	Id       		 	int    `json:"user_id"`
	Bio_discription  	string `json:"Bio"`
	Linked_in        	string `json:"linked in"`
	Instagram        	string `json:"instagram"`
	Github           	string `json:"github"`
	Facebook           	string `json:"facebook"`
	User			 	string `json:"user"` //forieng key reference to user
}

type OrdinaryUser struct {
	Id                  int    `json:"user_id"`
	Proffession         string `json:"Proffesion"`
	Participated_events string `json:"participated_events"`
	User                string `json:"user"` //forieng key reference to user
}


type Organization struct {
	Id       			int    `json:"user_id"`
	Location 			string `json:"location"`
	Organized_events 	string `json:"organized_events"`
	User     			string `json:"user"` //forieng key reference to user
}


type Event struct {
	gorm.Model
	User                     string `json:"user"` //forieng key reference to user
	Title                    string `json:"title"`
	Date                     string	`json:"date"`
	Location                 string	`json:"location"`
	Public_status            string `json:"public_status"`
	Offline                  bool	`json:"offline"`
	Free                     bool   `json:"rree"`
	Short_description        string `json:"short_description"`
	Long_description         string `json:"long_description"`
	Application_link         string `json:"application_link"`
	Website_link             string `json:"website_link"`
	Limit_applications_Id    string `json:"limit_applications"`
	Application_closing_date string `json:"application_closing_date"`
	Sub_events               string	`json:"sub_events"`
	Archived                 bool   `json:"archived"`
	Event_pic                string `json:"event_pic"`
}

type Bookmarks struct {
	gorm.Model
	User_Id 				string `json:"user_id"` //forieng key reference to user
	Event_Id 				string `json:"event_id"` //forieng key reference to event
}

type Applied_events struct {
	gorm.Model
	User_Id 				string `json:"user_id"` //forieng key reference to user
	Event_Id 				string `json:"event_id"` //forieng key reference to event
	Application_status 		string `json:"application_status"`
	Participation_status 	bool 	`json:"participation_status"`
}


type Posted_events struct {
	gorm.Model
	User_Id 				string `json:"user_id"` //forieng key reference to user
	Event_Id 				string `json:"event_id"` //forieng key reference to event
	Application_count 		int 	`json:"application_count"`
}

type Admin struct {
	ID       				int    `json:"id" `
	Username 				string `json:"username" gorm:"primary_key"`
	Password 				string `json:"password"`
}

//to store mail verification details

type Verification struct {
	gorm.Model
	Email 			string 	`json:"email"`
	Code  			int    	`json:"code"`
}

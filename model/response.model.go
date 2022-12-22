package model

import "gorm.io/gorm"

type UserResponse struct {
	gorm.Model

	ID              int    `json:"user_id"`
	First_Name      string `json:"first_name"`
	Last_Name       string `json:"last_name"`
  	Username		string `json:"username"`
	Email           string `json:"email"`
  	Phone           int64  `json:"phone_number"`
	Password        string `json:"password,omitempty"`
	Verified        bool   `json:"verified"`
	Profile         string `json:"profile"`
	Token			string `json:"token"`


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


type AdminResponse struct {
	ID       				int    `json:"id" `
	Username 				string `json:"username"`
	Password 				string `json:"password"`
	Token					string `json:"token"`
}

package interfaces

import (
	"radar/pkg/model"
	"radar/pkg/utils"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (model.UserResponse, error)
	AllUsers(pagenation utils.Filter) ([]model.UserResponse, utils.Metadata, error)
	InsertUser(user model.User) (int, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	CreateEvent(event model.Event) (string, error)
	FilterEventsBy(sex string, cusat_only string, free string) ([]model.EventResponse, error)
	AllEvents() ([]model.EventResponse, error)
	AskQuestion(question model.FAQA) error
	GetFaqa(event_name string) ([]model.FAQAResponse, error)
	GetQuestions(event_name string) ([]model.FAQAResponse, error)
	Answer(faqa model.FAQA, id string) error
	PostedEvents(organizer_name string) ([]model.EventResponse, error)
	UpdateUserinfo(user model.User, username string) error
	UpdatePassword(user model.User, email string, username string) error
	DeleteEvent(title string) error
}

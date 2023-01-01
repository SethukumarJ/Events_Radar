package interfaces

import "radar/pkg/model"

type UserService interface {
	CreateUser(user model.User) error
	FindUser(email string) (*model.UserResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
	CreateEvent(newEvent model.Event) (string, error)
	FilterEventsBy(sex string, cusat_only string, free string) (*[]model.EventResponse, error)
	AllEvents() (*[]model.EventResponse, error)
	AskQuestion(newQuestion model.FAQA) error
	GetFaqa(event_name string) (*[]model.FAQAResponse, error)
	GetQuestions(event_name string) (*[]model.FAQAResponse, error)
	Answer(faqa model.FAQA, id string) error
	PostedEvents(organizer_name string) (*[]model.EventResponse, error)
	UpdateUserinfo(user model.User, username string) error
	UpdatePassword(user model.User, email string, username string) error
	DeleteEvent(title string) error
}

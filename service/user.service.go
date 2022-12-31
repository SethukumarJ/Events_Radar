package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"radar/config"
	"radar/model"
	"radar/repo"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user model.User) error
	FindUser(email string) (*model.UserResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
	CreateEvent(newEvent model.Event) (string, error)
	FilterEventsBy(sex string,cusat_only string, free string) (*[]model.EventResponse,error)
	GetEventByTitle(title string) (*[]model.EventResponse,error)
	AllEvents() (*[]model.EventResponse, error)
	AskQuestion(newQuestion model.FAQA) error
	GetFaqa(event_name string) (*[]model.FAQAResponse, error)
	GetQuestions(event_name string) (*[]model.FAQAResponse, error)
	Answer(faqa model.FAQA, id string) error
	PostedEvents(organizer_name string) (*[]model.EventResponse,error)
	UpdateUserinfo(user model.User ,username string) error
	UpdatePassword(user model.User ,email string, username string) error
	DeleteEvent(title string) error

}

type userService struct {
	userRepo   repo.UserRepository
	adminRepo  repo.AdminRepository
	mailConfig config.MailConfig
}

func NewUserService(
	userRepo repo.UserRepository,
	adminRepo repo.AdminRepository,
	mailConfig config.MailConfig) UserService {
	return &userService{
		userRepo:   userRepo,
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
	}
}

/// CreateUser creates the user
func (c *userService) CreateUser(newUser model.User) error {

	fmt.Println("create user from service")
	_, err := c.userRepo.FindUser(newUser.Email)
	fmt.Println("fund user", err)

	if err == nil {
		return errors.New("Username already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	newUser.Password = HashPassword(newUser.Password)
	fmt.Println("password", newUser.Password)
	_, err = c.userRepo.InsertUser(newUser)
	if err != nil {
		return err
	}
	return nil

}

// FindUser finds the user
func (c *userService) FindUser(email string) (*model.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// SendVerificationEmail sends the verification email

func (c *userService) SendVerificationEmail(email string) error {
	//to generate random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)

	fmt.Println("code: ", code)

	message := fmt.Sprintf(
		"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for usingEvents.\n with regards Team Events radar.",
		code,
	)

	// send random code to user's email
	if err := c.mailConfig.SendMail(email, message); err != nil {
		return err
	}
	fmt.Println("email sent: ", email)

	err := c.userRepo.StoreVerificationDetails(email, code)

	if err != nil {
		return err
	}

	return nil
}

// VerifyAccount verifies the account
func (c *userService) VerifyAccount(email string, code int) error {

	err := c.userRepo.VerifyAccount(email, code)

	if err != nil {
		return err
	}
	return nil
}

func (c *userService) CreateEvent(newEvent model.Event) (string, error) {
	_, err := c.userRepo.CreateEvent(newEvent)
	if err != nil {
		return newEvent.Title, err
	}
	return newEvent.Title, nil
}


func (c *userService) FilterEventsBy(sex string,cusat_only string, free string) (*[]model.EventResponse,error) {

	events, err := c.userRepo.FilterEventsBy( sex, cusat_only,free)
	// log.Println("metadata from service", metadata)
	if err != nil {
		return nil, err
	}

	return &events, nil
}

func (c *userService) GetEventByTitle(title string) (*[]model.EventResponse,error) {

	events, err := c.userRepo.GetEventByTitle(title)
	// log.Println("metadata from service", metadata)
	if err != nil {
		return nil, err
	}

	return &events, nil
}


func (c *userService) AllEvents() (*[]model.EventResponse,error) {

	events, err := c.userRepo.AllEvents()
	// log.Println("metadata from service", metadata)
	if err != nil {
		return nil, err
	}

	return &events, nil
}


func (c *userService) AskQuestion(newQuestion model.FAQA) error {
	err := c.userRepo.AskQuestion(newQuestion)
	if err != nil {
		return err
	}
	return nil
}

// FindUser finds the user
func (c *userService) GetFaqa(event_name string) (*[]model.FAQAResponse, error) {
	faqa, err := c.userRepo.GetFaqa(event_name)

	if err != nil {
		return nil, err
	}

	return &faqa, nil
}


func (c *userService) GetQuestions(event_name string) (*[]model.FAQAResponse, error) {
	faqa, err := c.userRepo.GetQuestions(event_name)

	if err != nil {
		return nil, err
	}

	return &faqa, nil
}


func (c *userService) Answer(faqa model.FAQA ,id string) error{
	c.userRepo.Answer(faqa,id)
	
	return nil


}


func (c *userService) PostedEvents(organizer_name string) (*[]model.EventResponse,error) {

	events, err := c.userRepo.PostedEvents(organizer_name)
	// log.Println("metadata from service", metadata)
	if err != nil {
		return nil, err
	}

	return &events, nil
}


func (c *userService) UpdateUserinfo(user model.User ,username string) error{
	c.userRepo.UpdateUserinfo(user, username)
	fmt.Println("user frm update userinfo services:",user)
	
	return nil


}
func (c *userService) UpdatePassword(user model.User ,email string, username string) error{


	user.Password = HashPassword(user.Password)
	fmt.Println("password", user.Password)
	c.userRepo.UpdatePassword(user, email, username)
	
	return nil


}

func (c *userService) DeleteEvent(title string) error{
	c.userRepo.DeleteEvent(title)
	
	return nil


}
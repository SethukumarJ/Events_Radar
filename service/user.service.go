package service

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"radar/config"
	"radar/model"
	"radar/repo"
	"time"
)

type UserService interface {
	CreateUser(user model.User) error
	FindUser(email string) (*model.UserResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
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

// CreateUser creates the user
func (c *userService) CreateUser(newUser model.User) error {

	fmt.Println("create user from service")
	_, err := c.userRepo.FindUser(newUser.Email)
	fmt.Println("fund user",err )
	
	if err == nil {
		return errors.New("Username already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	newUser.Password = HashPassword(newUser.Password)
	fmt.Println("password",newUser.Password)
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
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
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


func (c *userService) CreateEvent(newEvent model.Event) (int, error) {
	_, err := c.userRepo.CreateEvent(newEvent)
	if err != nil {
		return newEvent.ID, err
	}
	return newEvent.ID, nil
}
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
	FindUser(username string) (*model.UserResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
}

type userService struct {
	userRepo   repo.UserRepository
	mailConfig config.MailConfig
}

func NewUserService(userRepo repo.UserRepository,
	mailConfig config.MailConfig) UserService {
	return &userService{
		userRepo:   userRepo,
		mailConfig: mailConfig,
	}
}

// CreateUser creates the user
func (u *userService) CreateUser(user model.User) error {

	_, err := u.userRepo.FindUser(user.Username)

	if err == nil {
		return errors.New("User already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	user.Password = HashPassword(user.Password)
	_, err = u.userRepo.InsertUser(&user)
	if err != nil {
		return err
	}
	return nil
}

// FindUser finds the user
func (c *userService) FindUser(username string) (*model.UserResponse, error) {
	user, err := c.userRepo.FindUser(username)

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
func (u *userService) SendVerificationEmail(email string) error {
	// Generate a random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)

	message := fmt.Sprintf(
		"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for usingEvents.\n with regards Team Events radar.",
		code,
	)

	// Send the email
	if err := u.mailConfig.SendMail(email, message); err != nil {
		return err
	}
	fmt.Println("Email sent successfully to: ",email)
	
	// Save the code in the database
	if err := u.userRepo.StoreVerificationDetails(email, code); 
	err != nil {
		return err
	}

	return nil
}


// VerifyAccount verifies the account
func (u *userService) VerifyAccount(email string, code int) error {
	err := u.userRepo.VerifiyAccount(email, code)
	if err != nil {
		return err
	}

	return nil
}

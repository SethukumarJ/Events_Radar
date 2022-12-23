package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"radar/repo"
)

// AuthService is the interface for authentication service
type AuthService interface {
	VerifyAdmin(email string, password string) error
	VerifyUser(email string, password string) error
}

// authService is the struct for authentication service
type authService struct {
	adminRepo repo.AdminRepository
	userRepo  repo.UserRepository
}

// NewAuthService returns a new instance of AuthService
func NewAuthService(
	adminRepo repo.AdminRepository,
	userRepo repo.UserRepository,

) AuthService {
	return &authService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}

// VerifyAdmin verifies the admin credentials
func (c *authService) VerifyAdmin(email, password string) error {

	admin, err := c.adminRepo.FindAdmin(email)

	//_, err = c.adminRepo.FindAdmin(email)

	if err != nil {
		return errors.New("Invalid Username/ password, failed to login")
	}

	isValidPassword := VerifyPassword(password, admin.Password)
	if !isValidPassword {
		return errors.New("Invalid username/ Password, failed to login")
	}

	return nil
}

// VerifyUser verifies the user credentials
func (c *authService) VerifyUser(email string, password string) error {

	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(password, user.Password)
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

// VerifyPassword verifies the password
func VerifyPassword(requestPassword, dbPassword string) bool {

	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword
}

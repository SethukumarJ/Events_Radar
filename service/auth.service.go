package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"radar/repo"
)

// AuthService is the interface for authentication service
type AuthService interface {
	VerifyAdmin(username, password string) (bool, error)
	VerifyUser(username, password string) (bool, error)
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
func (a *authService) VerifyAdmin(username, password string) (bool, error) {
	admin, err := a.adminRepo.FindAdmin(username)
	if err != nil {
		return false, errors.New("admin not found")
	}

	if !VerifyPassword(password, admin.Password) {
		return false, errors.New("invalid password")
	}

	return true, nil

}

// VerifyUser verifies the user credentials
func (a *authService) VerifyUser(username, password string) (bool, error) {
	user, err := a.userRepo.FindUser(username)
	if err != nil {
		return false, errors.New("user not found")
	}

	if !VerifyPassword(password, user.Password) {
		return false, errors.New("invalid password")
	}

	return true, nil
}

// VerifyPassword verifies the password
func VerifyPassword(requestPassword, dbPassword string) bool {

	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword
}

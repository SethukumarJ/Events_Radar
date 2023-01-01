package services

import (
	"errors"
	"fmt"
	"log"
	repository "radar/pkg/repository/interface"
	services "radar/pkg/services/interface"

	"golang.org/x/crypto/bcrypt"
)

// authService is the struct for authentication service
type authService struct {
	adminRepo repository.AdminRepository
	userRepo  repository.UserRepository
}

// NewAuthService returns a new instance of AuthService
func NewAuthService(
	adminRepo repository.AdminRepository,
	userRepo repository.UserRepository,

) services.AuthService {
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

	fmt.Println("adminpassword", admin.Password)
	fmt.Println("password:", password)

	isValidPassword := VerifyPassword(admin.Password, []byte(password))
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

	isValidPassword := VerifyPassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

// // VerifyPassword verifies the password
// func VerifyPassword(requestPassword, dbPassword string) bool {

// 	fmt.Println(requestPassword)
// 	requestPassword = HashPassword(requestPassword)
// 	fmt.Println(requestPassword)
// 	return requestPassword == dbPassword
// }

func VerifyPassword(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

package handler

import (
	"fmt"
	"log"
	"net/http"
	"radar/pkg/common/response"
	"radar/pkg/model"
	service "radar/pkg/services/interface"
	"radar/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthHandler interface
// type AuthHandler interface {
// 	AdminSignup()
// 	AdminLogin()
// 	UserSignup()
// 	UserLogin()
// 	AdminRefreshToken()
// 	UserRefreshToken()
// 	GoogleSignin()
// }

type AuthHandler struct {
	jwtAdminService service.JWTService
	jwtUserService  service.JWTService
	authService     service.AuthService
	adminService    service.AdminService
	userService     service.UserService
}

// NewAuthHandler returns a new instance of AuthHandler
func NewAuthHandler(
	jwtAdminService service.JWTService,
	jwtUserService service.JWTService,
	authService service.AuthService,
	adminService service.AdminService,
	userService service.UserService,
) AuthHandler {
	return AuthHandler{
		jwtAdminService: jwtAdminService,
		jwtUserService:  jwtUserService,
		authService:     authService,
		adminService:    adminService,
		userService:     userService,
	}
}

// AdminSignup handles the admin signup
func (cr *AuthHandler) AdminSignup(c *gin.Context) {

	var newAdmin model.Admin

	//fetching data
	c.Bind(&newAdmin)

	err := cr.adminService.CreateAdmin(newAdmin)

	if err != nil {
		response := response.ErrorResponse("Failed to signup", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	admin, _ := cr.adminService.FindAdmin(newAdmin.Username)
	admin.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", admin)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// AdminLogin handles the admin login
func (cr *AuthHandler) AdminLogin(c *gin.Context) {

	var adminLogin model.Admin

	c.Bind(&adminLogin)

	fmt.Println("adminLogin.passwrodk", adminLogin.Password)
	fmt.Println("adminLogin.username", adminLogin.Username)
	//verify User details
	err := cr.authService.VerifyAdmin(adminLogin.Username, adminLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	admin, _ := cr.adminService.FindAdmin(adminLogin.Username)
	token := cr.jwtUserService.GenerateToken(admin.ID, admin.Username, "admin")
	admin.Password = ""
	admin.Token = token
	response := response.SuccessResponse(true, "SUCCESS", admin.Token)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// admin refresh token
func (cr *AuthHandler) AdminRefreshToken(c *gin.Context) {

	// autheader := r.Header.Get("Authorization")
	autheader := c.Writer.Header().Get("Authorization")
	bearerToken := strings.Split(autheader, " ")
	token := bearerToken[1]

	refreshToken, err := cr.jwtAdminService.GenerateRefreshToken(token)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", refreshToken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// UserSignup handles the user signup

func (cr *AuthHandler) UserSignup(c *gin.Context) {

	var newUser model.User
	fmt.Println("user signup")
	//fetching data
	c.Bind(&newUser)

	//check username exit or not

	err := cr.userService.CreateUser(newUser)

	log.Println(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.userService.FindUser(newUser.Email)
	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// UserLogin handles the user login

func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var userLogin model.User

	c.Bind(&userLogin)

	//verify User details
	err := cr.authService.VerifyUser(userLogin.Email, userLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	user, _ := cr.userService.FindUser(userLogin.Email)
	token := cr.jwtUserService.GenerateToken(user.ID, user.Email, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user.Token)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// user refresh token
func (cr *AuthHandler) UserRefreshToken(c *gin.Context) {

	autheader := ("Authorization")
	bearerToken := strings.Split(autheader, " ")
	token := bearerToken[1]

	refreshToken, err := cr.jwtUserService.GenerateRefreshToken(token)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", refreshToken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"radar/common/response"
	"radar/model"
	"radar/service"
	"radar/utils"
	"strings"
)

// AuthHandler interface
type AuthHandler interface {
	AdminSignup() http.HandlerFunc
	AdminLogin() http.HandlerFunc
	UserSignup() http.HandlerFunc
	UserLogin() http.HandlerFunc
	AdminRefreshToken() http.HandlerFunc
	UserRefreshToken() http.HandlerFunc
}

type authHandler struct {
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
	return &authHandler{
		jwtAdminService: jwtAdminService,
		jwtUserService:  jwtUserService,
		authService:     authService,
		adminService:    adminService,
		userService:     userService,
	}
}

// AdminSignup handles the admin signup
func (c *authHandler) AdminSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newAdmin model.Admin

		//fetching data
		json.NewDecoder(r.Body).Decode(&newAdmin)

		err := c.adminService.CreateAdmin(newAdmin)

		if err != nil {
			response := response.ErrorResponse("Failed to signup", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		admin, _ := c.adminService.FindAdmin(newAdmin.Username)
		admin.Password = ""
		response := response.SuccessResponse(true, "SUCCESS", admin)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

// AdminLogin handles the admin login
func (c *authHandler) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var adminLogin model.Admin

		json.NewDecoder(r.Body).Decode(&adminLogin)

		//verifying  admin credentials
		err := c.authService.VerifyAdmin(adminLogin.Username, adminLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, response)
			return
		}

		//getting admin values
		admin, _ := c.adminService.FindAdmin(adminLogin.Username)
		token := c.jwtAdminService.GenerateToken(admin.ID, admin.Username, "admin")
		admin.Password = ""
		admin.Token = token
		response := response.SuccessResponse(true, "SUCCESS", admin.Token)
		utils.ResponseJSON(w, response)
	}

}

// admin refresh token
func (c *authHandler) AdminRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		autheader := r.Header.Get("Authorization")
		bearerToken := strings.Split(autheader, " ")
		token := bearerToken[1]

		refreshToken, err := c.jwtAdminService.GenerateRefreshToken(token)

		if err != nil {
			response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "SUCCESS", refreshToken)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}
// UserSignup handles the user signup

func (c *authHandler) UserSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newUser model.User
		fmt.Println("user signup")
		//fetching data
		json.NewDecoder(r.Body).Decode(&newUser)

		//check username exit or not

		err := c.userService.CreateUser(newUser)

		log.Println(newUser)

		if err != nil {
			response := response.ErrorResponse("Failed to create user", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		user, _ := c.userService.FindUser(newUser.Email)
		user.Password = ""
		response := response.SuccessResponse(true, "SUCCESS", user)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

// UserLogin handles the user login

func (c *authHandler) UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userLogin model.User

		json.NewDecoder(r.Body).Decode(&userLogin)

		//verify User details
		err := c.authService.VerifyUser(userLogin.Email, userLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, response)
			return
		}

		//fetching user details
		user, _ := c.userService.FindUser(userLogin.Email)
		token := c.jwtUserService.GenerateToken(user.ID, user.Email, "user")
		user.Password = ""
		user.Token = token
		response := response.SuccessResponse(true, "SUCCESS", user.Token)
		utils.ResponseJSON(w, response)

		fmt.Println("login function returned successfully")
	}

}

// user refresh token
func (c *authHandler) UserRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		autheader := r.Header.Get("Authorization")
		bearerToken := strings.Split(autheader, " ")
		token := bearerToken[1]

		refreshToken, err := c.jwtUserService.GenerateRefreshToken(token)

		if err != nil {
			response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "SUCCESS", refreshToken)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}


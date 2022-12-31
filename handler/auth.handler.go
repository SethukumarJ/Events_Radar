package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"radar/common/response"
	"radar/model"
	"radar/service"
	"radar/utils"
	"strings"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/subosito/gotenv"
)

// AuthHandler interface
type AuthHandler interface {
	AdminSignup() http.HandlerFunc
	AdminLogin() http.HandlerFunc
	UserSignup() http.HandlerFunc
	UserLogin() http.HandlerFunc
	AdminRefreshToken() http.HandlerFunc
	UserRefreshToken() http.HandlerFunc
	GoogleSignin()
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

		fmt.Println("adminLogin.passwrodk", adminLogin.Password)
		fmt.Println("adminLogin.username", adminLogin.Username)
		//verify User details
		err := c.authService.VerifyAdmin(adminLogin.Username, adminLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(w, response)
			return
		}

		//fetching user details
		admin, _ := c.adminService.FindAdmin(adminLogin.Username)
		token := c.jwtUserService.GenerateToken(admin.ID, admin.Username, "admin")
		admin.Password = ""
		admin.Token = token
		response := response.SuccessResponse(true, "SUCCESS", admin.Token)
		utils.ResponseJSON(w, response)

		fmt.Println("login function returned successfully")
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

func init() {
	err := gotenv.Load()

	if err != nil {
		log.Fatal("error loafding env")
	}

}

func (u *authHandler) GoogleSignin() {

	CLIENT_ID := os.Getenv("CLIENT_ID")         //get client id from env
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET") //get client secret key from env
	SECRET_KEY := os.Getenv("SECRET_KEY")       // get secret key for session handling
	REDIRECT_URL := os.Getenv("REDIRECT_URL")   //redirect url

	fmt.Println(CLIENT_ID)
	fmt.Println(CLIENT_SECRET)
	fmt.Println(SECRET_KEY)
	fmt.Println(REDIRECT_URL)

	// gothic.Store = store

	goth.UseProviders(
		google.New(CLIENT_ID, CLIENT_SECRET, REDIRECT_URL, "email", "profile"),
	)

	p := pat.New()
	p.Get("/auth/{provider}/callback", GetUser)

	p.Get("/auth/{provider}", AuthBigginer)

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("Templates/index.html")
		t.Execute(res, false)
	})

}

func GetUser(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	fmt.Println(user.Name, user.Email, user.FirstName, user.LastName)

	fmt.Println(user)
	// t, _ := template.ParseFiles("Templates/success.html")
	// t.Execute(res, user)
}

func AuthBigginer(res http.ResponseWriter, req *http.Request) {

	gothic.BeginAuthHandler(res, req)
}

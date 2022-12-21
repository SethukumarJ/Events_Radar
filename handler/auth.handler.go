package handler

import (
	"encoding/json"
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
func (h *authHandler) AdminSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newAdmin model.Admin

		//fetching the data from the request
		json.NewDecoder(r.Body).Decode(&newAdmin)

		err := h.adminService.CreateAdmin(newAdmin)

		if err != nil {
			response := response.ErrorResponse("Failed to create admin", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
	}

}

// AdminLogin handles the admin login
func (h *authHandler) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var adminLogin model.Admin

		//fetching the data from the request
		json.NewDecoder(r.Body).Decode(&adminLogin)

		//verifying the admin
		_, err := h.authService.VerifyAdmin(adminLogin.Username, adminLogin.Password)

		if err != nil {
			response := response.ErrorResponse("Failed to login", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		//getting admin values
		admin, err := h.adminService.FindAdmin(adminLogin.Username)
		token := h.jwtAdminService.GenerateToken(admin.ID, admin.Username, "admin")
		admin.Password = ""
		admin.Token = token
		response := response.SuccessResponse(true, "Login successful", admin.Token)
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

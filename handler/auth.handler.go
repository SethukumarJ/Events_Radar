package handler

import (
	"encoding/json"
	"net/http"
	"radar/common/response"
	"radar/model"
	"radar/service"
	"radar/utils"
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

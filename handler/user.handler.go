package handler

import (
	"net/http"
	"radar/common/response"
	"radar/service"
	"radar/utils"
	"strconv"
)

type UserHandler interface {
	SendVerificationMail() http.HandlerFunc
	VerifyAccount() http.HandlerFunc

}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// SendVerificationEmail sends the verification email

func (h *userHandler) SendVerificationMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := r.URL.Query().Get("Username")
		email := r.URL.Query().Get("Email")

		_, err := h.userService.FindUser(username)
		if err == nil {
			err = h.userService.SendVerificationEmail(email)
		}

		if err != nil {
			response := response.ErrorResponse("Failed to send verification email", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "Verification email sent", nil)
		utils.ResponseJSON(w, response)

	}
}

// verifyAccount verifies the account

func (c *userHandler) VerifyAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")
		code, _ := strconv.Atoi(r.URL.Query().Get("Code"))

		err := c.userService.VerifyAccount(email, code)

		if err != nil {
			response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Account verified successfully", email)
		utils.ResponseJSON(w, response)
	}
}
package handler

import (
	"fmt"
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

func (c *userHandler) SendVerificationMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("Email")

		_, err := c.userService.FindUser(email)
		fmt.Println("email: ", email)
		fmt.Println("err: ", err)

		if err == nil {
			err = c.userService.SendVerificationEmail(email)
		}

		fmt.Println(err)

		if err != nil {
			response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Verification mail sent successfully", email)
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



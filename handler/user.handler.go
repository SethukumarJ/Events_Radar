package handler

import (
	"net/http"
	"radar/common/response"
	"radar/service"
	"radar/utils"
)

type UserHandler interface {
	SendVerificationEmail() http.HandlerFunc
	VerifyEmail() http.HandlerFunc
}

type userHandler struct {
	userService service.UserService
}

func newUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// SendVerificationEmail sends the verification email

func (h *userHandler) SendVerificationMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := r.URL.Query().Get("email")

		_, err := h.userService.FindUser(email)
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


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
	"strconv"
)

type UserHandler interface {
	SendVerificationMail() http.HandlerFunc
	VerifyAccount() http.HandlerFunc
	CreateEvent() http.HandlerFunc
	FilterEventsBy() http.HandlerFunc
	AllEvents()  http.HandlerFunc
	AskQuestion() http.HandlerFunc
	GetFaqa() http.HandlerFunc
	GetQuestions() http.HandlerFunc
	Answer() http.HandlerFunc
	PostedEvents() http.HandlerFunc
	UpdateUserinfo() http.HandlerFunc
	UpdatePassword() http.HandlerFunc 
	DeleteEvent() http.HandlerFunc 
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

func (c *userHandler) CreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newEvent model.Event
		json.NewDecoder(r.Body).Decode(&newEvent)
		newEvent.Organizer_name = (r.Header.Get("Organizer_name"))
		_, err := c.userService.CreateEvent(newEvent)
		if err != nil {
			response := response.ErrorResponse("Failed to add new post", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", newEvent)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}

func (c *userHandler) FilterEventsBy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		free := r.URL.Query().Get("Free")
		sex := r.URL.Query().Get("Sex")
		fmt.Println("free from handlers:",free)
		cusat_only := (r.URL.Query().Get("Cusat_only"))
		fmt.Println("cusat only from handler:",cusat_only)

		events, err := c.userService.FilterEventsBy( sex,cusat_only, free)

		result := struct {
			Events *[]model.EventResponse
		}{
			Events: events,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All Events", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}


func (c *userHandler) AllEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		events, err := c.userService.AllEvents()

		result := struct {
			Events *[]model.EventResponse
		}{
			Events: events,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All Events", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}


func (c *userHandler) AskQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newQuestion model.FAQA
		json.NewDecoder(r.Body).Decode(&newQuestion)

		if newQuestion.Question == "" {

				log.Fatal("Qustion box is empty!")
				return
		}
		newQuestion.Username = r.Header.Get("User_name")
		newQuestion.Event_name = r.Header.Get("Event_name")
		err := c.userService.AskQuestion(newQuestion)
		if err != nil {
			response := response.ErrorResponse("Failed to add new comment", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", newQuestion)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
} 


func (c *userHandler) GetFaqa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		event_name := r.URL.Query().Get("Event_name")

		faqas, err := c.userService.GetFaqa(event_name)

		result := struct {
			Faqas *[]model.FAQAResponse
		}{
			Faqas: faqas,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All Events", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

		
	}
}

func (c *userHandler) GetQuestions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		event_name := r.URL.Query().Get("Event_name")

		faqas, err := c.userService.GetQuestions(event_name)

		result := struct {
			Faqas *[]model.FAQAResponse
		}{
			Faqas: faqas,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All Events", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

		
	}
}



func (c *userHandler) Answer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAnswer model.FAQA
		json.NewDecoder(r.Body).Decode(&newAnswer)
		newAnswer.Event_name = (r.Header.Get("Event_name"))
		id := r.URL.Query().Get("id")
		 err := c.userService.Answer(newAnswer, id)
		if err != nil {
			response := response.ErrorResponse("Failed to add new answer", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", newAnswer)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}


func (c *userHandler) PostedEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user_name := r.URL.Query().Get("Username")
	
		fmt.Println("free from handlers:",user_name)
	

		events, err := c.userService.PostedEvents(user_name)

		result := struct {
			Events *[]model.EventResponse
		}{
			Events: events,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "All Events", result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)

	}
}


func (c *userHandler) UpdateUserinfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateUser model.User
		json.NewDecoder(r.Body).Decode(&updateUser)
		
		username := r.URL.Query().Get("username")
		 err := c.userService.UpdateUserinfo(updateUser, username)
		if err != nil {
			response := response.ErrorResponse("Failed to apdate user", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", updateUser)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}


func (c *userHandler) UpdatePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatePassword model.User
		json.NewDecoder(r.Body).Decode(&updatePassword)
		
		username := r.URL.Query().Get("username")
		email := r.URL.Query().Get("email")
		 err := c.userService.UpdatePassword(updatePassword, email, username)
		if err != nil {
			response := response.ErrorResponse("Failed to apdate password", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "SUCCESS", updatePassword)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, response)
	}
}


func (c *userHandler) DeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("Title")
		

		err := c.userService.DeleteEvent(title)

		if err != nil {
			response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Deleted event successfully!", title)
		utils.ResponseJSON(w, response)
	}
}

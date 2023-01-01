package handler

import (
	"fmt"
	"log"
	"net/http"
	"radar/pkg/common/response"
	"radar/pkg/model"
	service "radar/pkg/services/interface"
	"radar/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type UserHandler interface {
// 	SendVerificationMail()
// 	VerifyAccount()
// 	CreateEvent()
// 	FilterEventsBy()
// 	AllEvents()
// 	AskQuestion()
// 	GetFaqa()
// 	GetQuestions()
// 	Answer()
// 	PostedEvents()
// 	UpdateUserinfo()
// 	UpdatePassword()
// 	DeleteEvent()
// }

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler{
	return UserHandler{
		userService: userService,
	}
}

// SendVerificationEmail sends the verification email

func (cr *UserHandler) SendVerificationMail(c *gin.Context) {

	email := c.Query("Email")

	_, err := cr.userService.FindUser(email)
	fmt.Println("email: ", email)
	fmt.Println("err: ", err)

	if err == nil {
		err = cr.userService.SendVerificationEmail(email)
	}

	fmt.Println(err)

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Verification mail sent successfully", email)
	utils.ResponseJSON(*c, response)

}

// verifyAccount verifies the account

func (cr *UserHandler) VerifyAccount(c *gin.Context) {

	email := c.Query("Email")
	code, _ := strconv.Atoi(c.Query("Code"))

	err := cr.userService.VerifyAccount(email, code)

	if err != nil {
		response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Account verified successfully", email)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) CreateEvent(c *gin.Context) {

	var newEvent model.Event
	c.Bind(&newEvent)
	newEvent.Organizer_name = (c.Writer.Header().Get("Organizer_name"))
	_, err := cr.userService.CreateEvent(newEvent)
	if err != nil {
		response := response.ErrorResponse("Failed to add new post", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", newEvent)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) FilterEventsBy(c *gin.Context) {

	free := c.Query("Free")
	sex := c.Query("Sex")
	fmt.Println("free from handlers:", free)
	cusat_only := (c.Query("Cusat_only"))
	fmt.Println("cusat only from handler:", cusat_only)

	events, err := cr.userService.FilterEventsBy(sex, cusat_only, free)

	result := struct {
		Events *[]model.EventResponse
	}{
		Events: events,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "All Events", result)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) AllEvents(c *gin.Context) {

	events, err := cr.userService.AllEvents()

	result := struct {
		Events *[]model.EventResponse
	}{
		Events: events,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "All Events", result)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) AskQuestion(c *gin.Context) {

	var newQuestion model.FAQA
	c.Bind(&newQuestion)

	if newQuestion.Question == "" {

		log.Fatal("Qustion box is empty!")
		return
	}
	newQuestion.Username = c.Writer.Header().Get("User_name")
	newQuestion.Event_name = c.Writer.Header().Get("Event_name")
	err := cr.userService.AskQuestion(newQuestion)
	if err != nil {
		response := response.ErrorResponse("Failed to add new comment", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", newQuestion)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) GetFaqa(c *gin.Context) {

	event_name := c.Query("Event_name")

	faqas, err := cr.userService.GetFaqa(event_name)

	result := struct {
		Faqas *[]model.FAQAResponse
	}{
		Faqas: faqas,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "All Events", result)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) GetQuestions(c *gin.Context) {

	event_name := c.Query("Event_name")

	faqas, err := cr.userService.GetQuestions(event_name)

	result := struct {
		Faqas *[]model.FAQAResponse
	}{
		Faqas: faqas,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "All Events", result)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) Answer(c *gin.Context) {

	var newAnswer model.FAQA
	c.Bind(&newAnswer)
	newAnswer.Event_name = (c.Writer.Header().Get("Event_name"))
	id := c.Query("id")
	err := cr.userService.Answer(newAnswer, id)
	if err != nil {
		response := response.ErrorResponse("Failed to add new answer", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", newAnswer)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) PostedEvents(c *gin.Context) {

	user_name := c.Query("Username")

	fmt.Println("free from handlers:", user_name)

	events, err := cr.userService.PostedEvents(user_name)

	result := struct {
		Events *[]model.EventResponse
	}{
		Events: events,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting posts from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "All Events", result)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) UpdateUserinfo(c *gin.Context) {

	var updateUser model.User
	c.Bind(&updateUser)

	username := c.Query("username")
	err := cr.userService.UpdateUserinfo(updateUser, username)
	if err != nil {
		response := response.ErrorResponse("Failed to apdate user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", updateUser)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) UpdatePassword(c *gin.Context) {

	var updatePassword model.User
	c.Bind(&updatePassword)

	username := c.Query("username")
	email := c.Query("email")
	err := cr.userService.UpdatePassword(updatePassword, email, username)
	if err != nil {
		response := response.ErrorResponse("Failed to apdate password", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", updatePassword)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *UserHandler) DeleteEvent(c *gin.Context) {

	title := c.Query("Title")

	err := cr.userService.DeleteEvent(title)

	if err != nil {
		response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Deleted event successfully!", title)
	utils.ResponseJSON(*c, response)

}

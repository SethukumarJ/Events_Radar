package handler

import (
	"fmt"
	"log"
	"net/http"
	"radar/common/response"
	"radar/model"
	"radar/service"
	"radar/utils"
	"strconv"
)

type AdminHandler interface {
	ApproveEvent() http.HandlerFunc
	ViewAllUsers() http.HandlerFunc
	ViewAllEventsFromAdminPanel() http.HandlerFunc
}

type adminHandler struct {
	adminService service.AdminService
	userService  service.UserService
}

func NewAdminHandler(
	adminService service.AdminService,
	userService service.UserService,

) AdminHandler {
	return &adminHandler{
		adminService: adminService,
		userService:  userService,
	}
}

func (c *adminHandler) ApproveEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("Title")

		err := c.adminService.ApproveEvent(title)

		if err != nil {
			response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}
		response := response.SuccessResponse(true, "Account verified successfully", title)
		utils.ResponseJSON(w, response)
	}
}

func (c *adminHandler) ViewAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))

		log.Println(page, "   ", pageSize)

		fmt.Println("page :", page)
		fmt.Println("pagesize", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		fmt.Println("pagenation",pagenation)

		users, metadata, err := c.adminService.AllUsers(pagenation)

		fmt.Println("users:",users)

		result := struct {
			Users *[]model.UserResponse
			Meta  *utils.Metadata
		}{
			Users: users,
			Meta:  metadata,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "Listed All Users", result)
		utils.ResponseJSON(w, response)

	}
}


func (c *adminHandler) ViewAllEventsFromAdminPanel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pagesize"))

		log.Println(page, "   ", pageSize)

		fmt.Println("page :", page)
		fmt.Println("pagesize", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		approved := r.URL.Query().Get("approved")
		fmt.Println("pagenation",pagenation)

		events, metadata, err := c.adminService.AllEventsInAdminPanel(pagenation,approved)

		fmt.Println("events:",events)

		result := struct {
			Events *[]model.EventResponse
			Meta  *utils.Metadata
		}{
			Events: events,
			Meta:  metadata,
		}

		if err != nil {
			response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(w, response)
			return
		}

		response := response.SuccessResponse(true, "Listed All Evnts", result)
		utils.ResponseJSON(w, response)

	}
}

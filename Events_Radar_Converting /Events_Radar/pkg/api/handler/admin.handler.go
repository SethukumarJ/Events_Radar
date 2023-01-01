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

// type AdminHandler interface {
// 	ApproveEvent() 
// 	ViewAllUsers() 
// 	ViewAllEventsFromAdminPanel() 
// }

type AdminHandler struct {
	adminService service.AdminService
	userService  service.UserService
}

func NewAdminHandler(
	adminService service.AdminService,
	userService service.UserService,

) AdminHandler {
	return AdminHandler{
		adminService: adminService,
		userService:  userService,
	}
}

func (cr *AdminHandler) ApproveEvent(c *gin.Context)  {
	
		title := c.Query("Title")
	
		err := cr.adminService.ApproveEvent(title)

		if err != nil {
			response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(*c, response)
			return
		}
		response := response.SuccessResponse(true, "Account verified successfully", title)
		utils.ResponseJSON(*c, response)
	
}

func (cr *AdminHandler) ViewAllUsers(c *gin.Context){
	

		page, _ := strconv.Atoi(c.Query("page"))

		pageSize, _ := strconv.Atoi(c.Query("pagesize"))

		log.Println(page, "   ", pageSize)

		fmt.Println("page :", page)
		fmt.Println("pagesize", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		fmt.Println("pagenation",pagenation)

		users, metadata, err := cr.adminService.AllUsers(pagenation)

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
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(*c, response)
			return
		}

		response := response.SuccessResponse(true, "Listed All Users", result)
		utils.ResponseJSON(*c, response)

	
}


func (cr *AdminHandler) ViewAllEventsFromAdminPanel(c *gin.Context) {


		page, _ := strconv.Atoi(c.Query("page"))

		pageSize, _ := strconv.Atoi(c.Query("pagesize"))

		log.Println(page, "   ", pageSize)

		fmt.Println("page :", page)
		fmt.Println("pagesize", pageSize)

		pagenation := utils.Filter{
			Page:     page,
			PageSize: pageSize,
		}

		approved := c.Query("approved")
		fmt.Println("pagenation",pagenation)

		events, metadata, err := cr.adminService.AllEventsInAdminPanel(pagenation,approved)

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
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusBadRequest)
			utils.ResponseJSON(*c, response)
			return
		}

		response := response.SuccessResponse(true, "Listed All Evnts", result)
		utils.ResponseJSON(*c, response)

	
}

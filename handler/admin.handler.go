package handler

import "radar/service"

type AdminHandler interface {
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

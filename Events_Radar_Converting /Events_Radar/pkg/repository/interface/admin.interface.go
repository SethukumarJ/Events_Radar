package interfaces

import (
	"radar/pkg/model"
	"radar/pkg/utils"
)

type AdminRepository interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (model.AdminResponse, error)
	ApproveEvent(title string) error
	AllEventsInAdminPanel(pagenation utils.Filter, approved string) ([]model.EventResponse, utils.Metadata, error)
}

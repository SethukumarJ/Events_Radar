package service

import (
	"database/sql"
	"errors"
	"log"
	"radar/model"
	"radar/repo"
)


// AdminService is the interface for admin service
type AdminService interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (*model.AdminResponse, error)
}

// adminService is the struct for admin service
type adminService struct {
	adminRepo repo.AdminRepository
	userRepo  repo.UserRepository
}

// NewAdminService returns a new instance of AdminService
func NewAdminService(
	adminRepo repo.AdminRepository,
	userRepo repo.UserRepository,
) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}

// FindAdmin finds the admin
func (c *adminService) FindAdmin(username string) (*model.AdminResponse, error) {
	admin, err := c.adminRepo.FindAdmin(username)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}


// CreateAdmin creates the admin
func (c *adminService) CreateAdmin(admin model.Admin) error {

	_, err := c.adminRepo.FindAdmin(admin.Username)

	if err == nil {
		return errors.New("Admin already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	admin.Password = HashPassword(admin.Password)
	err = c.adminRepo.CreateAdmin(admin)

	if err != nil {
		log.Println(err)
		return errors.New("error while signup")
	}
	return nil
}

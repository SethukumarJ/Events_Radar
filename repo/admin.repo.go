package repo

import (
	"database/sql"
	"radar/model"
)

type AdminRepository interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (model.AdminResponse, error)
}

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}
}

func (a *adminRepo) CreateAdmin(admin model.Admin)error {
	var err error
	query := `INSERT INTO admins (
					username,
					password,
			) VALUES ($1,$2);`

	 err = a.db.QueryRow(query,
		admin.Username,
		admin.Password,
	).Err()

	return err

}

func (a *adminRepo) FindAdmin(username string) (model.AdminResponse, error) {
	var admin model.AdminResponse
	query := `SELECT id,
			 		username,
					password
					FROM admins WHERE username=$1;`
	err := a.db.QueryRow(query, username).Scan(
		&admin.ID,
		&admin.Username,
		&admin.Password,
	)
	return admin, err
}
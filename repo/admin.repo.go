package repo

import (
	"database/sql"
	"errors"
	"log"
	"radar/model"
)

type AdminRepository interface {
	CreateAdmin(admin model.Admin) error
	FindAdmin(username string) (model.AdminResponse, error)
	ApproveEvent(title string) error
}

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}
}

func (c *adminRepo) CreateAdmin(admin model.Admin) error {

	query := `INSERT INTO
				admins (username,password)
				VALUES
				($1, $2);`
	err := c.db.QueryRow(
		query, admin.Username,
		admin.Password,
	).Err()
	return err
}

func (c *adminRepo) FindAdmin(username string) (model.AdminResponse, error) {

	log.Println("username of admin:", username)
	var admin model.AdminResponse

	query := `SELECT
			id, 
			username,
			password
			FROM admins WHERE username = $1;`

	err := c.db.QueryRow(query,
		username).Scan(
		&admin.ID,
		&admin.Username,
		&admin.Password)

	return admin, err
}

func (c *adminRepo) ApproveEvent(title string) error {

	var id int

	query := `SELECT id FROM 
				events WHERE 
				title = $1;`
	err := c.db.QueryRow(query, title).Scan(&id)

	if err == sql.ErrNoRows {
		return errors.New("invalid title")
	}

	if err != nil {
		return err
	}

	query = `UPDATE events SET
				approved = $1
				WHERE
				title = $2 ;`
	err = c.db.QueryRow(query, true, title).Err()
	log.Println("Updating approval status to true ", err)
	if err != nil {
		return err
	}

	return nil
}

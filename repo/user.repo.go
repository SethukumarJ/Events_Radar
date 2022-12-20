package repo

import (
	"database/sql"
	"radar/model"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	InsertUser(user *model.User) (int, error)
}

// UserRepo is a struct that represent the UserRepo's repository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo will create an object that represent the UserRepo's repository interface
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// InsertUser will create a new user
func (u *UserRepo) InsertUser(user *model.User) (string, error) {
	var err error
	var username string
	query := `INSERT INTO users (
		Id,
		first_name,
		last_name,
		email,
		phone,
		username,
		password,
		profile,
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	RETURNING id`
	err = u.db.QueryRow(
		query,
		user.Id,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Username,
		user.Phone,
		user.Password,
		user.Profile,
		).Scan(&username)
	
	return username, err
}

package repo

import (
	"database/sql"
	"radar/model"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	InsertUser(user *model.User) (int, error)
	FindUser(id int) (*model.User, error)
}

// UserRepo is a struct that represent the UserRepo's repository
type userRepo struct {
	db *sql.DB
}

// NewUserRepo will create an object that represent the UserRepo's repository interface
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

// InsertUser will create a new user
func (u *userRepo) InsertUser(user *model.User) (string, error) {
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
		user.Phone,
		user.Username,
		user.Password,
		user.Profile,
		).Scan(&username)
	
	return username, err
}

// FindUser will return a user with a given email
func(u *userRepo)  FindUser(email string) (*model.User, error) {
	var err error
	var user model.User
	query := `SELECT * FROM users WHERE username = $1`
	err = u.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Phone,
		&user.Username,
		&user.Password,
		&user.Profile,
	)
	return &user, err
}


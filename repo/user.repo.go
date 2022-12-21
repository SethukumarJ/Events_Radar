package repo

import (
	"database/sql"
	"log"
	"radar/model"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	InsertUser(user *model.User) (string, error)
	FindUser(email string) (model.UserResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifiyAccount(email string, code int) error
}

// UserRepo is a struct that represent the UserRepo's repository
type userRepo struct {
	db *sql.DB
}

// NewUserRepo will create an object that represent the UserRepo's repository interface
func NewUserRepo(db *sql.DB) UserRepository {
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
	RETURNING username;`
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
func (u *userRepo) FindUser(email string) (model.UserResponse, error) {
	var err error
	var user model.UserResponse
	query := `SELECT * FROM users WHERE username = $1;`
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
	return user, err
}


// StoreVerificationDetails will store the verification details

func (u *userRepo) StoreVerificationDetails(email string, code int) error {

	var err error
	query := `INSERT INTO 
				verification (email, code) VALUES 
				($1, $2);`

	err = u.db.QueryRow(query, email, code).Err()
	return err
}


// VerifiyAccount will verify the user account

func (u *userRepo) VerifiyAccount(email string, code int) error {

	var id int
	var err error

	query := `SELECT id FROM 
				verification WHERE 
				email = $1 AND code = $2;`
	err = u.db.QueryRow(query, email, code).Scan(&id)

	if err != nil {
		return err
	}

	query = `UPDATE users SET
				verification = true WHERE
				email = $1;`

	err = u.db.QueryRow(query, email).Err()
	log.Println("Updating User verification status", err)
	if err != nil {
		return err
	}

	return nil

}

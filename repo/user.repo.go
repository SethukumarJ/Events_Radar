package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"radar/model"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (model.UserResponse, error)
	InsertUser(user model.User) (int, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	CreateEvent(event model.Event) (string, error)
	AllEvents() ([]model.EventResponse, error)
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
func (c *userRepo) InsertUser(user model.User) (int, error) {

	var id int

	query := `INSERT INTO users(
			first_name,
			last_name,
			email,
			phone,
			password,
			profile)
			VALUES
			($1, $2, $3, $4, $5, $6)
			RETURNING id;`

	err := c.db.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Phone,
		user.Password,
		user.Profile).Scan(
		&id,
	)

	fmt.Println("id", id)
	return id, err
}

// FindUser will return a user with a given email
func (c *userRepo) FindUser(email string) (model.UserResponse, error) {

	var user model.UserResponse

	query := `SELECT 
				id,
				first_name,
				last_name,
				email,
				password,
				phone,
				profile
				FROM users 
				WHERE email = $1;`

	err := c.db.QueryRow(query,
		email).Scan(
		&user.ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Profile,
	)

	fmt.Println("user from find user :", user)
	return user, err
}

// StoreVerificationDetails will store the verification details

func (u *userRepo) StoreVerificationDetails(email string, code int) error {

	var err error
	query := `INSERT INTO 
				verifications (email, code) VALUES 
				($1, $2);`

	err = u.db.QueryRow(query, email, code).Err()
	return err
}

// VerifiyAccount will verify the user account

func (c *userRepo) VerifyAccount(email string, code int) error {

	var id int

	query := `SELECT id FROM 
				verifications WHERE 
				email = $1 AND code = $2;`
	err := c.db.QueryRow(query, email, code).Scan(&id)

	if err == sql.ErrNoRows {
		return errors.New("Invalid verification code/Email")
	}

	if err != nil {
		return err
	}

	query = `UPDATE users SET
				verification = $1
				WHERE
				email = $2 ;`
	err = c.db.QueryRow(query, true, email).Err()
	log.Println("Updating User verification: ", err)
	if err != nil {
		return err
	}

	return nil
}

func (c *userRepo) CreateEvent(event model.Event) (string, error) {
	var title string

	query := `INSERT INTO events(
		created_at,
		organizer_name,
		title,
		event_date,
		location,
		offline,
		cusat_only,
		Free,
		short_description,
		long_description,
		application_link,
		website_link,
		application_closing_date,
		sub_events,
		event_pic
		)VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
			RETURNING title;`

	err := c.db.QueryRow(query,
		event.Created_at,
		event.Organizer_name,
		event.Title,
		event.Event_date,
		event.Location,
		event.Offline,
		event.Cusat_only,
		event.Free,
		event.Short_description,
		event.Long_description,
		event.Application_link,
		event.Website_link,
		event.Application_closing_date,
		event.Sub_events,
		event.Event_pic).Scan(
		&title,
	)

	fmt.Println(title)
	return title, err

}

func (c *userRepo) AllEvents() ([]model.EventResponse, error){


	var events []model.EventResponse

	query := `SELECT 
				COUNT(*) OVER(),
				created_at,
				organizer_name,
				title,
				event_pic,
				event_date,
				location,
				offline,
				short_description,
				long_description,
				application_link,
				website_link,
				application_closing_date,
				sub_events,
				free
				FROM events WHERE cusat_only = $1;`

				rows, err := c.db.Query(query, true)

				if err != nil {
					return nil, err
				} 

				var totalRecords int

				defer rows.Close()


				for rows.Next() {
					var Event model.EventResponse
			
					err = rows.Scan(
						&totalRecords,
						&Event.Created_at,
						&Event.Organizer_name,
						&Event.Title,
						&Event.Event_pic,
						&Event.Event_date,
						&Event.Location,
						&Event.Offline,
						&Event.Short_description,
						&Event.Long_description,
						&Event.Application_link,
						&Event.Website_link,
						&Event.Application_closing_date,
						&Event.Sub_events,
						&Event.Free,

						)
			
					if err != nil {
						return events,  err
					}
					events = append(events, Event)
				}

			log.Println(events)

			return events, nil


	}

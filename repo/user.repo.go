package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"radar/model"
	"radar/utils"
	"reflect"
	// "github.com/pelletier/go-toml/query"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (model.UserResponse, error)
	AllUsers(pagenation utils.Filter) ([]model.UserResponse, utils.Metadata, error)
	InsertUser(user model.User) (int, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	CreateEvent(event model.Event) (string, error)
	FilterEventsBy(sex string, cusat_only string, free string) ([]model.EventResponse, error)
	AllEvents() ([]model.EventResponse, error)
	AskQuestion(question model.FAQA) error
	GetFaqa(event_name string) ([]model.FAQAResponse, error)
	GetQuestions(event_name string) ([]model.FAQAResponse, error)
	Answer(faqa model.FAQA, id string) error 
	PostedEvents(organizer_name string) ([]model.EventResponse, error)
	UpdateUserinfo(user model.User, username string) (error)
	UpdatePassword(user model.User,email string, username string) error
	DeleteEvent(title string) error
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
			username,
			first_name,
			last_name,
			email,
			phone_number,
			password,
			profile)
			VALUES
			($1, $2, $3, $4, $5, $6,$7)
			RETURNING id;`

	err := c.db.QueryRow(query,
		user.Username,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Phone_number,
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
				phone_number,
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

	query = `DELETE FROM verifications WHERE email = $1;`

	err = c.db.QueryRow(query,email).Err()
	fmt.Println("deleting the verification code.")
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
		event_pic,
		sex
		)VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
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
		event.Event_pic,
		event.Sex).Scan(
		&title,
	)

	fmt.Println(title)
	return title, err

}

func (c *userRepo) FilterEventsBy(sex string, cusat_only string, free string) ([]model.EventResponse, error) {

	var events []model.EventResponse

	fmt.Println(free, "from repo")
	fmt.Println(sex, "sexfrom repo")
	fmt.Println("cusat only from repo:", cusat_only)

	fmt.Println(reflect.TypeOf(sex))
	fmt.Println(reflect.TypeOf(free))

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
				FROM events WHERE approved = true AND cusat_only = $1 AND sex = $2 AND free = $3;`

	rows, err := c.db.Query(query, cusat_only, sex, free)

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
			return events, err
		}
		events = append(events, Event)
	}

	log.Println(events)

	return events, nil

}

func (c *userRepo) AllEvents() ([]model.EventResponse, error) {

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
				FROM events WHERE approved = true;`

	rows, err := c.db.Query(query)

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
			return events, err
		}
		events = append(events, Event)
	}

	log.Println(events)

	return events, nil

}

func (c *userRepo) AskQuestion(question model.FAQA) error {

	query := `INSERT INTO faqas(
		created_at,
		question,
		event_name,
		username)VALUES
			($1, $2, $3, $4);`
	err := c.db.QueryRow(query,
		question.CreatedAt,
		question.Question,
		question.Event_name,
		question.Username)

		fmt.Println("question :",question.Question)
	log.Println("error : ", err)
	if err == nil {
		return errors.New("Failed to post queston!")
	}
	return nil
}

func (c *userRepo) GetFaqa(event_name string) ([]model.FAQAResponse, error) {

	var faqas []model.FAQAResponse
	fmt.Println("Getquestion called")
	query := `SELECT 
	COUNT(*) OVER(),
	id,
	event_name,
	created_at,
	username,
	question,
	answer
	FROM faqas WHERE public = $1 AND event_name = $2;`

	rows, err := c.db.Query(query,true,event_name)

	fmt.Println("rows",rows)

	if err != nil {
		return nil, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var Faqas model.FAQAResponse

		err = rows.Scan(
			&totalRecords,
			&Faqas.Id,
			&Faqas.Event_name,
			&Faqas.CreatedAt,
			&Faqas.Username,
			&Faqas.Question,
			&Faqas.Answer,
		)

		if err != nil {
			return faqas, err
		}
		faqas = append(faqas, Faqas)
	}

	fmt.Println(faqas)
	log.Println(faqas)

	return faqas, nil

}



func (c *userRepo) GetQuestions(event_name string) ([]model.FAQAResponse, error) {

	var faqas []model.FAQAResponse
	fmt.Println("Getquestion called")
	query := `SELECT 
	COUNT(*) OVER(),
	id,
	event_name,
	created_at,
	username,
	question
	FROM faqas WHERE public = $1 AND event_name = $2;`

	rows, err := c.db.Query(query,false,event_name)

	fmt.Println("rows",rows)

	if err != nil {
		return nil, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var Faqas model.FAQAResponse

		err = rows.Scan(
			&totalRecords,
			&Faqas.Id,
			&Faqas.Event_name,
			&Faqas.CreatedAt,
			&Faqas.Username,
			&Faqas.Question,
		)

		if err != nil {
			return faqas, err
		}
		faqas = append(faqas, Faqas)
	}

	fmt.Println(faqas)
	log.Println(faqas)

	return faqas, nil

}

func (c *userRepo) Answer(faqa model.FAQA,id string) error {



	query := `UPDATE faqas SET
	public = true,
	answer = $1
	WHERE id = $2 AND
	event_name = $3;`
	err := c.db.QueryRow(query, faqa.Answer,id,faqa.Event_name).Err()
	log.Println("Updating faqas answer: ", err)
	if err != nil {
		return err
	}

	return nil

}

func (c *userRepo) AllUsers(pagenation utils.Filter) ([]model.UserResponse, utils.Metadata, error) {

	fmt.Println("allusers called from repo")
	var users []model.UserResponse

	query := `SELECT 
				COUNT(*) OVER(),
				first_name,
				last_name,
				username,
				email,
				phone_number,
				profile
				FROM users
				LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("allusers called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allusers called from repo")

	for rows.Next() {
		var User model.UserResponse
		fmt.Println("username :", User.Username)
		err = rows.Scan(
			&totalRecords,
			&User.First_Name,
			&User.Last_Name,
			&User.Username,
			&User.Email,
			&User.Phone,
			&User.Profile,
		)

		fmt.Println("username", User.Username)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		users = append(users, User)
	}

	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

}



func (c *userRepo) PostedEvents(organizer_name string) ([]model.EventResponse, error) {

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
				FROM events WHERE approved = true AND organizer_name = $1;`

	rows, err := c.db.Query(query,organizer_name)

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
			return events, err
		}
		events = append(events, Event)
	}

	log.Println(events)

	return events, nil

}



func (c *userRepo) UpdateUserinfo(user model.User, username string) (error) {

	fmt.Println("phone",user.Phone_number)
	fmt.Println("email",user.Email)
	fmt.Println("profile",user.Profile)
	query := `Update users SET
			username = $1,
			email = $2,
			phone_number = $3,
			profile = $4
			WHERE username = $5`

	err := c.db.QueryRow(query,
		user.Username,
		user.Email,
		user.Phone_number,
		user.Profile,
		username).Err()
		
		fmt.Println("Updating userinfo: ", err)
		if err != nil {
			return err
		}
	
		return nil
	}


	func (c *userRepo) UpdatePassword(user model.User,email string, username string) (error) {


		query := `Update users SET
				password = $1,
				WEHRE email = $2 OR username = $3`
	
		err := c.db.QueryRow(query,
			user.Password,
			email,
			username).Err()
			
			log.Println("Updating userpassword: ", err)
			if err != nil {
				return err
			}
		
			return nil
		}


		func (c *userRepo) DeleteEvent(title string) error {

				query := `DELETE FROM events
						    WHERE title = $1;`

			    err := c.db.QueryRow(query,title).Err()
				log.Println("Deleted event : ", err)
					if err != nil {
				return err
				}
		
			return nil
		}
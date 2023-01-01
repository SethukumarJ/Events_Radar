package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"radar/pkg/model"
	repository "radar/pkg/repository/interface"
	"radar/pkg/utils"
)

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) repository.AdminRepository {
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

func (c *adminRepo) AllEventsInAdminPanel(pagenation utils.Filter, approved string) ([]model.EventResponse, utils.Metadata, error) {

	fmt.Println("allevents called from repo")
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
				FROM events WHERE approved = $1
				LIMIT $2 OFFSET $3;`

	rows, err := c.db.Query(query, approved, pagenation.Limit(), pagenation.Offset())
	fmt.Println("approved", approved)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("allusers called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("alllevents called from repo")

	for rows.Next() {
		var Event model.EventResponse
		fmt.Println("username :", Event.Title)
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
		fmt.Println("username", Event.Title)

		if err != nil {
			return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		events = append(events, Event)
	}

	if err := rows.Err(); err != nil {
		return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(events)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

}

// func (c *userRepo) AllUsers(pagenation utils.Filter) ([]model.UserResponse, utils.Metadata, error) {

// 	fmt.Println("allusers called from repo")
// 	var users []model.UserResponse

// 	query := `SELECT
// 				COUNT(*) OVER(),
// 				first_name,
// 				last_name,
// 				email,
// 				phone,
// 				profile
// 				FROM users
// 				LIMIT $1 OFFSET $2;`

// 	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())
// 	fmt.Println("rows", rows)
// 	if err != nil {
// 		return nil, utils.Metadata{}, err
// 	}

// 	fmt.Println("allusers called from repo")

// 	var totalRecords int

// 	defer rows.Close()
// 	fmt.Println("allusers called from repo")

// 	for rows.Next() {
// 		var User model.UserResponse
// 		fmt.Println("username :", User.Username)
// 		err = rows.Scan(
// 			&totalRecords,
// 			&User.First_Name,
// 			&User.Last_Name,
// 			&User.Email,
// 			&User.Phone,
// 			&User.Profile,
// 		)

// 		fmt.Println("username", User.Username)

// 		if err != nil {
// 			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
// 		}
// 		users = append(users, User)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
// 	}
// 	log.Println(users)
// 	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
// 	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

// }

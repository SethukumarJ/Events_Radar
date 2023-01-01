package db

import (
	"database/sql"
	"fmt"
	"log"
	"radar/pkg/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) *sql.DB {

	databaseName := cfg.DBName

	dbURI := cfg.DBSOURCE

	//open connection to database

	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	// verify connectioin to the database is still alive

	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging database")
		log.Fatal(err)
	}

	log.Println("/nConnected to database", databaseName)

	return db

}

package config

import (
	"fmt"
	"log"
	"os"
	"radar/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dbURL := os.Getenv("DB_SOURCE")
	fmt.Println("connected to:", dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Verification{})
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.FAQA{})

	// db.Exec(`INSERT INTO admins (
	// 			username,password)
	// 		VALUES (
	// 			$1,$2)`,
	// 	"admin@gmail.com", "admin")

	return db

}

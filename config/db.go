package config

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dbURL := os.Getenv("DB_SOURCE")
	db, err := gorm.Open(postges.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Verification{})

	// db.Exec(`INSERT INTO admins (
	// 			username,password)
	// 		VALUES (
	// 			$1,$2)`,
	// 	"admin@gmail.com", "admin")

	return db

}

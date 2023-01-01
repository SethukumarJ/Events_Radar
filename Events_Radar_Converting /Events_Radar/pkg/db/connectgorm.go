package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"radar/pkg/config"
	"radar/pkg/model"
)

func ConnectGormDB(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	psqlInfo := cfg.DBSOURCE
	fmt.Printf("\n\nsql : %v\n\n", psqlInfo)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Verification{})
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.FAQA{})

	return db, dbErr
}



































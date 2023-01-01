package main

import (
	"fmt"
	"log"

	
	_ "radar/pkg/common/response"
	_ "radar/pkg/model"

	"radar/pkg/config"
	"radar/pkg/db"

	di "radar/pkg/di"
)


func main() {
	fmt.Println("starting my job-portal project in clean code architecture")

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	// db.ConnectDB(config)
	gorm, _ := db.ConnectGormDB(config)
	fmt.Printf("\ngorm : %v\n\n", gorm)

	server, diErr := di.InitializeAPI(config)
	fmt.Printf("\n\n\nserver ; %v\n\n\n", server)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}

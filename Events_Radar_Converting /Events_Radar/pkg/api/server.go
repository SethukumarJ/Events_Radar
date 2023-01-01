package api

import (
	"fmt"
	"log"

	"radar/pkg/api/middleware"
	"radar/pkg/api/handler"

	"github.com/gin-gonic/gin"
	
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	authHandler *handler.AuthHandler,
	middleware middleware.Middleware,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handl
	// Request JWT

	//userroutes
	user := engine.Group("user")
	{
		user.POST("/signup", authHandler.UserSignup)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send/verification", userHandler.SendVerificationMail)
		user.PATCH("/verify/account", userHandler.VerifyAccount)
		user.POST("/CreateEvent", userHandler.CreateEvent)
		user.GET("/AllEvents", userHandler.AllEvents)
		user.GET("/FilterEventsBy", userHandler.FilterEventsBy)
		user.POST("/AskQuestion", userHandler.AskQuestion)
		user.GET("/GetFaqa", userHandler.GetFaqa)
		user.GET("/organizer/GetQuestions", userHandler.GetQuestions)
		user.PATCH("/organizer/Answer", userHandler.Answer)
		user.GET("/PostedEvents", userHandler.PostedEvents)
		user.PATCH("/updateinfo", userHandler.UpdateUserinfo)
		user.PATCH("/updatePassword", userHandler.UpdatePassword)
		user.DELETE("/event/delete", userHandler.DeleteEvent)
		
	user.Use(middleware.AuthorizeJwt())
		{
			user.GET("/token/refresh", authHandler.UserRefreshToken)
		}
	}


	//admin routes
	admin := engine.Group("admin")
{
	admin.POST("/signup", authHandler.AdminSignup)
	admin.POST("/login", authHandler.AdminLogin)
	admin.PATCH("/ApproveEvent",adminHandler.ApproveEvent)
	admin.GET("/view/users", adminHandler.ViewAllUsers)
	admin.GET("/view/Events", adminHandler.ViewAllEventsFromAdminPanel)
	
	user.Use(middleware.AuthorizeJwt())
	{
		user.GET("/token/refresh", authHandler.UserRefreshToken)
	}
}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	// fmt.Println("error at 8080:")
	fmt.Print("\n\nddddddddd\n\n")
	err := sh.engine.Run(":8080")
	fmt.Println("error at 8080:")
	if err != nil {
		log.Fatalln(err)
	}
}

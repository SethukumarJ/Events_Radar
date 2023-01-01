package diErr

import (
	http "radar/pkg/api"
	handler "radar/pkg/api/handler"
	middleware "radar/pkg/api/middleware"

	config "radar/pkg/config"
	"radar/pkg/db"
	repository "radar/pkg/repository"
	usecase "radar/pkg/services"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	// fmt.Printf("\n\n\nv\n\n\n")
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepo,
		repository.NewAdminRepo,
		config.NewMailConfig,
		config.LoadConfig,
		usecase.NewJWTUserService,
		usecase.NewJWTAdminService,
		usecase.NewUserService,
		usecase.NewAuthService,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewAuthHandler,
		middleware.NewMiddlewareUser,
		middleware.NewMiddlewareAdmin,
		http.NewServerHTTP)

	// fmt.Printf("\n\n\nbuild : %v\n\n\n", s)
	return &http.ServerHTTP{}, nil

}
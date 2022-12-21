package routes

import (
	h "radar/handler"
	m "radar/middleware"

	"github.com/go-chi/chi"
)

type UserRoute interface {
	UserRouter(routes chi.Router,
		authHandler h.AuthHandler,
		userHandler h.UserHandler,
		middleware m.Middleware)
}

type userRoute struct{}

func NewUserRoute() UserRoute {
	return &userRoute{}
}

func (r *userRoute) UserRouter(routes chi.Router,
	authHandler h.AuthHandler,
	userHandler h.UserHandler,
	middleware m.Middleware,
) {

	routes.Post("/user/signup", authHandler.UserSignup())
	routes.Post("/user/login", authHandler.UserLogin())
	routes.Post("/user/send/verification", userHandler.SendVerificationMail())
	routes.Patch("/user/verify/account", userHandler.VerifyAccount())

	routes.Group(func(r chi.Router) {
		r.Use(middleware.AuthorizeJwt)
		r.Get("/user/token/refresh", authHandler.UserRefreshToken())

	})

}

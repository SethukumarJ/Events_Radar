package routes

import (
	h "radar/handler"
	m "radar/middleware"

	"github.com/go-chi/chi"
)

type AdminRoute interface {
	AdminRouter(routes chi.Router,
		authHandler h.AuthHandler,
		adminHandler h.AdminHandler,
		middleware m.Middleware)
}

type adminRoute struct{}

func NewAdminRoute() AdminRoute {
	return &adminRoute{}
}

// to handle admin routes
func (r *adminRoute) AdminRouter(routes chi.Router,
	authHandler h.AuthHandler,
	adminHandler h.AdminHandler,
	middleware m.Middleware,
) {

	routes.Post("/admin/signup", authHandler.AdminSignup())
	routes.Post("/admin/login", authHandler.AdminLogin())

	routes.Group(func(r chi.Router) {
		r.Use(middleware.AuthorizeJwt)

		r.Get("/admin/token/refresh", authHandler.AdminRefreshToken())

	})

}

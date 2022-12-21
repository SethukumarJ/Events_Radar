package routes

import (
	"github.com/go-chi/chi"
	h "radar/handler"
	m "radar/middleware"
	)

type AdminRoute interface {
	AdminRouter(
		routes chi.Router,
		authHandler h.AuthHandler,
		adminHandler h.AdminHandler,
		middleware m.Middleware,
	)
}

package internal

import (
	openapi "github.com/mytord/fs/backend/gen/opencliapi"
	"github.com/mytord/fs/backend/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(publicRoutes openapi.Routes, privateRoutes openapi.Routes) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.Logger)

	//// public routes
	r.Group(func(r chi.Router) {
		applyRoutes(publicRoutes, r)
	})

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.Auth)
		applyRoutes(privateRoutes, r)
	})

	return r
}

func applyRoutes(routes openapi.Routes, router chi.Router) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.Method(route.Method, route.Pattern, handler)
	}
}

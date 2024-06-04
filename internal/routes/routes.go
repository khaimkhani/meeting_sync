package routes

import (
	"github.com/gorilla/mux"
	"msync/internal/handlers"
	"msync/internal/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello-world", handlers.PrimHandler)
	// TODO some routes need to allow not logged in users
	r.Use(middleware.AuthMiddleware)
}

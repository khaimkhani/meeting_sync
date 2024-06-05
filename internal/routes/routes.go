package routes

import (
	"github.com/gorilla/mux"
	"msync/internal/handlers"
	"msync/internal/middleware"
)

func RegisterRoutes(r *mux.Router) {
	// TODO some routes need to consider scope
	// either register subroutes and use perm middleware or
	// handle in handler with a helper funC
	r.HandleFunc("/hello-world", handlers.PrimHandler)
	r.Use(middleware.AuthMiddleware)
}

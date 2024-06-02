package routes

import (
	"github.com/gorilla/mux"
	"msync/internal/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello-world", handlers.PrimHandler)
}

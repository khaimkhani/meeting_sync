package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"msync/internal/routes"
	"net/http"
	"os"
)

func getDBUrl() string {
	return os.Getenv("DBSTRING")
}

var DB *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	// init db
	DB, err := sql.Open("postgres", getDBUrl())
	if err != nil {
		log.Fatal("Error connecting to postgres")
	}
	defer DB.Close()

	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	http.ListenAndServe(":5000", router)

}

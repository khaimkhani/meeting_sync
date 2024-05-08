package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func getDBUrl() string {
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	//
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v", dbUser, dbPwd, dbHost, dbPort, dbName)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "What's good")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
